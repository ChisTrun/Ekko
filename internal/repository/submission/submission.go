package submission

import (
	"context"
	ekko "ekko/api"
	"ekko/internal/utils/paging"
	utils "ekko/internal/utils/sort"
	"ekko/internal/utils/tx"
	"ekko/pkg/ent"
	"ekko/pkg/ent/question"
	"ekko/pkg/ent/scenario"
	"ekko/pkg/ent/scenariocandidate"
	"ekko/pkg/ent/submissionattempt"
	"errors"
	"fmt"
)

type Submission interface {
	Create(ctx context.Context, tx tx.Tx, candidateId uint64, req *ekko.SubmitAnswerRequest) (*ent.SubmissionAttempt, error)
	GetAttempt(ctx context.Context, user uint64, isBm bool, attemptId uint64) (*ent.SubmissionAttempt, error)
	ListAttempt(ctx context.Context, candidateId uint64, req *ekko.ListAttemptRequest) ([]*ent.SubmissionAttempt, int32, int32, error)
	ListAllSubmission(ctx context.Context, req *ekko.ListAllSubmissionRequest) ([]*ent.ScenarioCandidate, int32, int32, error)
}

type submission struct {
	ent *ent.Client
}

func New(ent *ent.Client) Submission {
	return &submission{
		ent: ent,
	}
}

func (s *submission) Create(ctx context.Context, tx tx.Tx, candidateId uint64, req *ekko.SubmitAnswerRequest) (*ent.SubmissionAttempt, error) {

	attemptNumber, err := s.getAvailableAttemptNumber(ctx, req.ScenarioId, candidateId)
	if err != nil {
		return nil, err
	}

	scenarioCandidate, err := tx.Client().ScenarioCandidate.Query().Where(scenariocandidate.CandidateID(candidateId), scenariocandidate.ScenarioID(req.ScenarioId)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if scenarioCandidate == nil {
		if scenarioCandidate, err = tx.Client().ScenarioCandidate.Create().SetCandidateID(candidateId).SetScenarioID(req.ScenarioId).Save(ctx); err != nil {
			return nil, err
		}
	}

	ok, err := s.checkSubmission(ctx, req.ScenarioId, req.Answers)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("invalid submission")
	}

	submission, err := tx.Client().SubmissionAttempt.Create().SetAttemptNumber(attemptNumber).SetScenarioCandidate(scenarioCandidate).Save(ctx)
	if err != nil {
		return nil, err
	}

	creates := []*ent.AnswerSubmissionCreate{}
	for _, question := range req.Answers {
		query := tx.Client().AnswerSubmission.Create().
			SetSubmissionAttemptID(submission.ID).
			SetQuestionID(question.QuestionId).
			SetAnswer(question.Answer).
			SetStatus(ekko.SubmissionStatus_SUBMISSION_STATUS_IN_PROGRESS)
		creates = append(creates, query)
	}
	answers, err := tx.Client().AnswerSubmission.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, err
	}

	submission.Edges.Answers = answers
	submission.Edges.ScenarioCandidate = scenarioCandidate

	if err := tx.Client().Scenario.UpdateOneID(req.ScenarioId).AddParticipants(1).Exec(ctx); err != nil {
		return nil, err
	}

	return submission, nil
}

func (s *submission) getAvailableAttemptNumber(ctx context.Context, scenarioId, candidateId uint64) (int32, error) {
	n, _ := s.ent.SubmissionAttempt.Query().Where(
		submissionattempt.HasScenarioCandidateWith(
			scenariocandidate.ScenarioID(scenarioId),
			scenariocandidate.CandidateID(candidateId)),
	).Aggregate(ent.Max(submissionattempt.FieldAttemptNumber)).
		Int(ctx)

	return int32(n + 1), nil
}

func (s *submission) checkSubmission(ctx context.Context, scenarioId uint64, answers []*ekko.SubmitAnswerRequest_SubmittedAnswer) (bool, error) {
	checkMap := make(map[uint64]bool)
	for _, answer := range answers {
		checkMap[answer.QuestionId] = true
	}

	questionIds, err := s.ent.Question.Query().Where(question.ScenarioID(scenarioId)).IDs(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return false, err
	}

	for _, id := range questionIds {
		if _, ok := checkMap[id]; !ok {
			return false, nil
		}
	}

	return true, nil
}

func (s *submission) GetAttempt(ctx context.Context, user uint64, isBm bool, attemptId uint64) (*ent.SubmissionAttempt, error) {
	query := s.ent.SubmissionAttempt.Query().Where(
		submissionattempt.ID(attemptId),
	).WithAnswers()

	if !isBm {
		query = query.Where(submissionattempt.HasScenarioCandidateWith(scenariocandidate.CandidateID(user)))
	} else {
		query = query.Where(submissionattempt.HasScenarioCandidateWith(scenariocandidate.HasScenarioWith(scenario.BmID(user))))
	}

	return query.Only(ctx)
}

func (s *submission) ListAttempt(ctx context.Context, candidateId uint64, req *ekko.ListAttemptRequest) ([]*ent.SubmissionAttempt, int32, int32, error) {
	query := s.ent.SubmissionAttempt.Query().Where(
		submissionattempt.HasScenarioCandidateWith(
			scenariocandidate.CandidateID(candidateId),
			scenariocandidate.ScenarioID(req.ScenarioId),
		),
	)

	sort, err := utils.GetSort(submissionattempt.Columns, submissionattempt.Table, req.SortMethod)
	if err != nil {
		return nil, 0, 0, err
	}

	totalCount, _ := query.Count(ctx)
	totalPage := paging.GetPagingData(int32(totalCount), req.PageSize)

	attempts, err := query.
		Modify(sort).
		Offset(int(req.PageIndex) * int(req.PageSize)).
		Limit(int(req.PageSize)).
		All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}
	return attempts, int32(totalCount), totalPage, nil
}

func (s *submission) ListAllSubmission(ctx context.Context, req *ekko.ListAllSubmissionRequest) ([]*ent.ScenarioCandidate, int32, int32, error) {
	query := s.ent.ScenarioCandidate.Query().Where(
		scenariocandidate.ScenarioID(req.ScenarioId),
	)

	if (req.From == nil && req.To != nil) || (req.From != nil && req.To == nil) {
		return nil, 0, 0, errors.New("invalid time")
	} else if req.From != nil {
		if req.From.AsTime().After(req.To.AsTime()) {
			return nil, 0, 0, errors.New("invalid time")
		}
		query = query.Where(scenariocandidate.CreatedAt(req.From.AsTime()), scenariocandidate.CreatedAt(req.To.AsTime()))
	}

	sort, err := utils.GetSort(submissionattempt.Columns, submissionattempt.Table, req.SortMethod)
	if err != nil {
		return nil, 0, 0, err
	}

	totalCount, _ := query.Count(ctx)
	totalPage := paging.GetPagingData(int32(totalCount), req.PageSize)

	attempts, err := query.
		Modify(sort).
		Offset(int(req.PageIndex) * int(req.PageSize)).
		Limit(int(req.PageSize)).
		WithAttempts().
		All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}
	return attempts, int32(totalCount), totalPage, nil
}
