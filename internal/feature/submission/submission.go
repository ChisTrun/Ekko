package submission

import (
	"context"
	ekko "ekko/api"
	"ekko/internal/repository"
	"ekko/internal/utils/checker"
	"ekko/internal/utils/converter"
	"ekko/internal/utils/extractor"
	"ekko/internal/utils/tx"
	bulbasaur "ekko/pkg/bulbasaur/api"
	"ekko/pkg/ent"
	"ekko/pkg/logger/pkg/logging"
	"fmt"
	"sync"
)

type Submission interface {
	ListAttempt(ctx context.Context, req *ekko.ListAttemptRequest) (*ekko.ListAttemptResponse, error)
	GetAttempt(ctx context.Context, req *ekko.GetAttemptRequest) (*ekko.GetAttemptResponse, error)
	SubmitAnswer(ctx context.Context, req *ekko.SubmitAnswerRequest) (*ekko.SubmitAnswerResponse, error)
	ListAllSubmission(ctx context.Context, req *ekko.ListAllSubmissionRequest) (*ekko.ListAllSubmissionResponse, error)
}

type submission struct {
	repo      *repository.Repository
	extractor extractor.Extractor
}

func New(repo *repository.Repository, extractor extractor.Extractor) Submission {
	return &submission{
		repo:      repo,
		extractor: extractor,
	}
}

func (s *submission) ListAttempt(ctx context.Context, req *ekko.ListAttemptRequest) (*ekko.ListAttemptResponse, error) {

	userId, err := s.extractor.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	attempts, totalCount, totalPage, err := s.repo.Submission.ListAttempt(ctx, uint64(userId), req)
	if err != nil {
		return nil, err
	}

	data := make([]*ekko.Attempt, len(attempts))
	var wg sync.WaitGroup

	for i, attempt := range attempts {
		wg.Add(1)
		go func(i int, attempt ent.SubmissionAttempt) {
			defer wg.Done()
			data[i] = converter.ConvertAttempt(&attempt)
		}(i, *attempt)
	}
	wg.Wait()

	return &ekko.ListAttemptResponse{
		Attempts:   data,
		TotalCount: totalCount,
		TotalPage:  totalPage,
		Request:    req,
	}, nil

}

func (s *submission) GetAttempt(ctx context.Context, req *ekko.GetAttemptRequest) (*ekko.GetAttemptResponse, error) {

	isBm := false
	roleIds := s.extractor.GetRoleIDs(ctx)
	if len(roleIds) != 1 {
		return nil, fmt.Errorf("invalid role")
	}

	switch roleIds[0] {
	case fmt.Sprintf("%v", int32(bulbasaur.Role_ROLE_BUSINESS_MANAGER)):
		isBm = true
	case fmt.Sprintf("%v", int32(bulbasaur.Role_ROLE_CANDIDATE)):
		isBm = false
	default:
		return nil, fmt.Errorf("invalid role")
	}

	userId, err := s.extractor.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	attempt, err := s.repo.Submission.GetAttempt(ctx, uint64(userId), isBm, req.Id)
	if err != nil {
		return nil, err
	}

	return &ekko.GetAttemptResponse{
		Attempt: converter.ConvertAttempt(attempt),
	}, nil
}

func (s *submission) SubmitAnswer(ctx context.Context, req *ekko.SubmitAnswerRequest) (*ekko.SubmitAnswerResponse, error) {

	userId, err := s.extractor.GetUserID(ctx)
	if err != nil {
		logging.Logger(ctx).Error(fmt.Sprintf("failed to get user id: %v", err))
		return nil, err
	}

	attempt := &ent.SubmissionAttempt{}
	if txErr := tx.WithTransaction(ctx, s.repo.Ent, func(ctx context.Context, tx tx.Tx) error {
		var err error
		attempt, err = s.repo.Submission.Create(ctx, tx, uint64(userId), req)
		return err
	}); txErr != nil {
		logging.Logger(ctx).Error(fmt.Sprintf("failed to submit answer: %v", txErr))
		return nil, txErr
	}

	return &ekko.SubmitAnswerResponse{
		Attempt: converter.ConvertAttempt(attempt),
	}, nil
}

func (s *submission) ListAllSubmission(ctx context.Context, req *ekko.ListAllSubmissionRequest) (*ekko.ListAllSubmissionResponse, error) {
	roleIds := s.extractor.GetRoleIDs(ctx)
	if err := checker.CheckRole(ctx, fmt.Sprintf("%v", int32(bulbasaur.Role_ROLE_BUSINESS_MANAGER)), roleIds); err != nil {
		return nil, err
	}

	scenarioCandidates, totalCount, totalPage, err := s.repo.Submission.ListAllSubmission(ctx, req)
	if err != nil {
		return nil, err
	}

	data := make([]*ekko.Submission, len(scenarioCandidates))
	var wg sync.WaitGroup

	for i, submission := range scenarioCandidates {
		wg.Add(1)
		go func(i int, submission ent.ScenarioCandidate) {
			defer wg.Done()
			data[i] = converter.ConvertSubmission(&submission)
		}(i, *submission)
	}
	wg.Wait()

	return &ekko.ListAllSubmissionResponse{
		Submissions: data,
		TotalCount:  totalCount,
		TotalPage:   totalPage,
		Request:     req,
	}, nil
}
