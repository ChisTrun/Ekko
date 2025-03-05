package question

import (
	"context"
	ekko "ekko/api"
	"ekko/internal/utils/tx"
	"ekko/pkg/ent"
	entquestion "ekko/pkg/ent/question"
)

type Question interface {
	CreateList(ctx context.Context, tx tx.Tx, scenarioId uint64, questions []*ekko.ScenarioQuestion) ([]*ent.Question, error)
}

type question struct {
}

func New() Question {
	return &question{}
}

func (q *question) CreateList(ctx context.Context, tx tx.Tx, scenarioId uint64, questions []*ekko.ScenarioQuestion) ([]*ent.Question, error) {
	tx.Client().Question.Delete().Where(entquestion.ScenarioID(scenarioId)).Exec(ctx)

	creates := []*ent.QuestionCreate{}
	for _, question := range questions {
		query := tx.Client().Question.Create().
			SetScenarioID(scenarioId).
			SetContent(question.Content).
			SetCriteria(question.Criteria).
			SetHint(question.Hint)
		creates = append(creates, query)
	}

	return tx.Client().Question.CreateBulk(creates...).Save(ctx)
}
