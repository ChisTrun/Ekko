package converter

import (
	ekko "ekko/api"
	"ekko/pkg/ent"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertField(ent *ent.ScenarioField) *ekko.Field {
	return &ekko.Field{
		Id:   ent.ID,
		Name: ent.Name,
		BaseData: &ekko.BaseData{
			CreatedAt: timestamppb.New(ent.CreatedAt),
			UpdatedAt: timestamppb.New(ent.UpdatedAt),
		},
	}
}

func ConvertScenario(ent *ent.Scenario, listing bool) *ekko.Scenario {
	out := &ekko.Scenario{
		Id:               ent.ID,
		Name:             ent.Name,
		Description:      ent.Description,
		BaseData:         &ekko.BaseData{CreatedAt: timestamppb.New(ent.CreatedAt), UpdatedAt: timestamppb.New(ent.UpdatedAt)},
		Rating:           float32(ent.Rating),
		Fields:           []*ekko.Field{},
		TotalParticipant: ent.Participants,
		Questions:        []*ekko.Question{},
		TotalQuestion:    int32(len(ent.Edges.Questions)),
	}

	for _, field := range ent.Edges.Field {
		out.Fields = append(out.Fields, ConvertField(field))
	}

	if !listing {
		for _, question := range ent.Edges.Questions {
			out.Questions = append(out.Questions, ConvertQuestion(question))
		}
	}

	return out
}

func ConvertQuestion(ent *ent.Question) *ekko.Question {
	return &ekko.Question{
		Id:       ent.ID,
		Criteria: ent.Criteria,
		Hint:     ent.Hint,
		Content:  ent.Content,
		BaseData: &ekko.BaseData{
			CreatedAt: timestamppb.New(ent.CreatedAt),
			UpdatedAt: timestamppb.New(ent.UpdatedAt),
		},
	}
}

func ConvertAttempt(ent *ent.SubmissionAttempt) *ekko.Attempt {
	out := &ekko.Attempt{
		Id:         ent.ID,
		ScenarioId: ent.Edges.ScenarioCandidate.ScenarioID,
		Answers:    []*ekko.Answer{},
		BaseData: &ekko.BaseData{
			CreatedAt: timestamppb.New(ent.CreatedAt),
			UpdatedAt: timestamppb.New(ent.UpdatedAt),
		},
		AttemptNumber: ent.AttemptNumber,
	}

	for _, answer := range ent.Edges.Answers {
		out.Answers = append(out.Answers, ConvertAnswer(answer))
	}

	return out
}

func ConvertAnswer(ent *ent.AnswerSubmission) *ekko.Answer {
	return &ekko.Answer{
		QuestionId:          ent.QuestionID,
		Answer:              ent.Answer,
		Relevance:           float32(ent.Relevance),
		ClarityCompleteness: float32(ent.ClarityCompleteness),
		Accuracy:            float32(ent.Accuracy),
		Overall:             float32(ent.Overall),
		Status:              ent.Status,
		BaseData: &ekko.BaseData{
			CreatedAt: timestamppb.New(ent.CreatedAt),
			UpdatedAt: timestamppb.New(ent.UpdatedAt),
		},
	}
}

func ConvertSubmission(ent *ent.ScenarioCandidate) *ekko.Submission {
	out := &ekko.Submission{
		Id:          ent.ID,
		ScenarioId:  ent.ScenarioID,
		CandidateId: ent.CandidateID,
		Attempts:    []*ekko.Attempt{},
	}

	for _, attempt := range ent.Edges.Attempts {
		out.Attempts = append(out.Attempts, ConvertAttempt(attempt))
	}

	return out
}
