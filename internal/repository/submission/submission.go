package submission

import (
	"context"
	ekko "ekko/api"
	"ekko/internal/rabbit"
	"ekko/internal/utils/paging"
	utils "ekko/internal/utils/sort"
	"ekko/internal/utils/tx"
	"ekko/pkg/ent"
	"ekko/pkg/ent/answersubmission"
	"ekko/pkg/ent/question"
	"ekko/pkg/ent/scenario"
	"ekko/pkg/ent/scenariocandidate"
	"ekko/pkg/ent/submissionattempt"
	"ekko/pkg/logger/pkg/logging"
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
)

type Submission interface {
	Create(ctx context.Context, tx tx.Tx, candidateId uint64, req *ekko.SubmitAnswerRequest) (*ent.SubmissionAttempt, error)
	GetAttempt(ctx context.Context, user uint64, isBm bool, attemptId uint64) (*ent.SubmissionAttempt, error)
	ListAttempt(ctx context.Context, candidateId uint64, req *ekko.ListAttemptRequest) ([]*ent.SubmissionAttempt, int32, int32, error)
	ListAllSubmission(ctx context.Context, req *ekko.ListAllSubmissionRequest) ([]*ent.ScenarioCandidate, int32, int32, error)
	ReceiveResponse(ctx context.Context, msg amqp.Delivery) error
	SendSubmission(ctx context.Context, scenario *ent.Scenario, attempId uint64)
}

type submission struct {
	ent      *ent.Client
	rabbitMQ rabbit.Rabbit
}

func New(ent *ent.Client, rabbitMQ rabbit.Rabbit) Submission {
	return &submission{
		ent:      ent,
		rabbitMQ: rabbitMQ,
	}
}

func (s *submission) Create(ctx context.Context, tx tx.Tx, candidateId uint64, req *ekko.SubmitAnswerRequest) (*ent.SubmissionAttempt, error) {

	// scenario, err := tx.Client().Scenario.Query().Where(scenario.ID(req.ScenarioId)).Select(scenario.FieldID, scenario.FieldName, scenario.FieldDescription).Only(ctx)
	// if err != nil {
	// 	return nil, err
	// }

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

	// s.sendSubmission(context.Background(), tx, scenario, submission.ID)
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
	).WithAnswers().WithScenarioCandidate()

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
		WithScenarioCandidate().
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
		query = query.Where(scenariocandidate.CreatedAtGTE(req.From.AsTime()), scenariocandidate.CreatedAtLTE(req.To.AsTime()))
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
		WithAttempts(func(saq *ent.SubmissionAttemptQuery) {
			saq.WithScenarioCandidate()
			saq.WithAnswers()
		}).
		All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}
	return attempts, int32(totalCount), totalPage, nil
}

func (s *submission) SendSubmission(ctx context.Context, scenario *ent.Scenario, attempId uint64) {

	answersubmissions, err := s.ent.AnswerSubmission.Query().
		Where(
			answersubmission.SubmissionAttemptID(attempId),
		).
		WithQuestion(func(qq *ent.QuestionQuery) {
			qq.Select(question.FieldContent, question.FieldCriteria)
		}).
		Select(answersubmission.FieldAnswer, answersubmission.FieldID).
		All(ctx)

	if err != nil {
		logging.Logger(ctx).Error("can not get answer submission", zap.Error(err))
	}

	evaluationReq := ekko.EvaluationRequest{
		Scenario: &ekko.EvaluationRequest_EvalutionScenario{
			Name:        scenario.Name,
			Description: scenario.Description,
		},
		Data: []*ekko.EvaluationRequest_QuestionAnswerPair{},
	}

	for _, answer := range answersubmissions {
		if answer.Edges.Question == nil {
			logging.Logger(ctx).Error("can not get question", zap.Error(err))
			return
		}
		evaluationReq.Data = append(evaluationReq.Data, &ekko.EvaluationRequest_QuestionAnswerPair{
			Id:       answer.ID,
			Question: answer.Edges.Question.Content,
			Answer:   answer.Answer,
			Criteria: answer.Edges.Question.Criteria,
		})
	}

	jsonData, err := protojson.Marshal(&evaluationReq)
	if err != nil {
		logging.Logger(ctx).Error("can not marshal evaluation request", zap.Error(err))
		return
	}

	if err := s.rabbitMQ.Publish(ctx, jsonData); err != nil {
		logging.Logger(ctx).Error("can not send message to rabbitmq", zap.Error(err))
		return
	}

}

func (s *submission) ReceiveResponse(ctx context.Context, msg amqp.Delivery) error {
	evaluationRsp := &ekko.EvaluationResponse{}
	if err := protojson.Unmarshal(msg.Body, evaluationRsp); err != nil {
		logging.Logger(ctx).Error("can not unmarshal evaluation response", zap.Error(err))
		return err
	}

	return tx.WithTransaction(ctx, s.ent, func(ctx context.Context, tx tx.Tx) error {
		for _, result := range evaluationRsp.Result {
			_, err := tx.Client().AnswerSubmission.UpdateOneID(result.Id).
				SetStatus(result.Status).
				SetAccuracy(result.Accuracy).
				SetClarityCompleteness(result.ClarityCompleteness).
				SetRelevance(result.Relevance).
				SetOverall(result.Overall).
				Save(ctx)
			if err != nil {
				logging.Logger(ctx).Error("can not update answer submission", zap.Error(err))
				return err
			}
		}
		return nil
	})
}
