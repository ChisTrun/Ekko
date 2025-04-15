package scenario

import (
	"context"
	ekko "ekko/api"
	"ekko/internal/repository/question"
	"ekko/internal/utils/paging"
	utils "ekko/internal/utils/sort"
	"ekko/internal/utils/tx"
	"ekko/pkg/ent"
	entquestion "ekko/pkg/ent/question"
	entscenario "ekko/pkg/ent/scenario"
	"ekko/pkg/ent/scenariocandidate"
	"ekko/pkg/ent/scenariofavorite"
	"ekko/pkg/ent/scenariofield"
	"errors"
	"fmt"
	"sync"
)

type Scenario interface {
	Create(ctx context.Context, tx tx.Tx, ownerId uint64, req *ekko.CreateScenarioRequest) (*ent.Scenario, error)
	Update(ctx context.Context, tx tx.Tx, bmId uint64, req *ekko.UpdateScenarioRequest) error
	Delete(ctx context.Context, tx tx.Tx, bmId uint64, ids []uint64) error
	List(ctx context.Context, req *ekko.ListScenarioRequest, userId *uint64) ([]*ent.Scenario, int32, int32, error)
	Get(ctx context.Context, req *ekko.GetScenarioRequest) (*ent.Scenario, error)
	Favorite(ctx context.Context, userId uint64, scenarioId uint64) error
	Rating(ctx context.Context, id uint64, rating float64) error
}

type scenario struct {
	ent      *ent.Client
	question question.Question
	mux      *sync.Mutex
}

func New(ent *ent.Client, question question.Question) Scenario {
	return &scenario{
		ent:      ent,
		question: question,
		mux:      &sync.Mutex{},
	}
}

func (s *scenario) Create(ctx context.Context, tx tx.Tx, ownerId uint64, req *ekko.CreateScenarioRequest) (*ent.Scenario, error) {
	scenario, err := tx.Client().Scenario.Create().
		SetName(req.Name).
		SetDescription(req.Description).
		SetDescription(req.Description).
		AddFieldIDs(req.FieldIds...).
		SetBmID(ownerId).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(req.Questions) != 0 {
		questions, err := s.question.CreateList(ctx, tx, scenario.ID, req.Questions)
		if err != nil {
			return nil, err
		}

		scenario.Edges.Questions = questions
	}

	return scenario, nil
}

func (s *scenario) Update(ctx context.Context, tx tx.Tx, bmId uint64, req *ekko.UpdateScenarioRequest) error {
	scenario, err := tx.Client().Scenario.Query().Where(
		entscenario.IDEQ(req.Id),
		entscenario.BmID(bmId),
	).
		WithCandidates().
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if len(scenario.Edges.Candidates) > 0 {
		return fmt.Errorf("could not update scenario with candidates")
	}

	query := scenario.Update()
	if len(req.FieldIds) > 0 {
		query = query.ClearFieldEdge().AddFieldIDs(req.FieldIds...)
	}

	if err := query.
		SetName(req.Name).
		SetDescription(req.Description).Exec(ctx); err != nil {
		return err
	}

	_, err = s.question.CreateList(ctx, tx, scenario.ID, req.Questions)

	return err
}

func (s *scenario) Delete(ctx context.Context, tx tx.Tx, bmId uint64, ids []uint64) error {
	_, err := tx.Client().Scenario.Delete().Where(entscenario.IDIn(ids...), entscenario.BmID(bmId)).Exec(ctx)
	return err
}

func (s *scenario) List(ctx context.Context, req *ekko.ListScenarioRequest, userId *uint64) ([]*ent.Scenario, int32, int32, error) {
	query := s.ent.Scenario.Query()

	if (req.From == nil && req.To != nil) || (req.From != nil && req.To == nil) {
		return nil, 0, 0, errors.New("invalid time")
	} else if req.From != nil {
		if req.From.AsTime().After(req.To.AsTime()) {
			return nil, 0, 0, errors.New("invalid time")
		}
		query = query.Where(entscenario.CreatedAtGTE(req.From.AsTime()), entscenario.CreatedAtLTE(req.To.AsTime()))
	}

	if len(req.BmIds) > 0 {
		query = query.Where(entscenario.BmIDIn(req.BmIds...))
	}

	if req.SearchContent != nil {
		query = query.Where(entscenario.Or(
			entscenario.NameContainsFold(*req.SearchContent),
			entscenario.DescriptionContainsFold(*req.SearchContent)))
	}

	if len(req.FieldIds) > 0 {
		query = query.Where(entscenario.HasFieldWith(scenariofield.IDIn(req.FieldIds...)))
	}

	if req.MinParticipant != nil {
		query = query.Where(entscenario.ParticipantsGTE(*req.MinParticipant))
	}

	if req.MinRating != nil {
		query = query.Where(entscenario.RatingGTE(float64(*req.MinRating)))
	}

	if userId != nil {
		if req.IsFavorite != nil && *req.IsFavorite {
			query = query.Where(entscenario.HasFavoritesWith(scenariofavorite.UserID(*userId)))
		}
		if req.IsFinished != nil && *req.IsFinished {
			query = query.Where(entscenario.HasCandidatesWith(scenariocandidate.CandidateID(*userId)))
		}
	}

	sort, err := utils.GetSort(entscenario.Columns, entscenario.Table, req.SortMethods)
	if err != nil {
		return nil, 0, 0, err
	}

	totalCount, _ := query.Count(ctx)
	totalPage := paging.GetPagingData(int32(totalCount), req.PageSize)

	scenarials, err := query.
		Modify(sort).
		Offset(int(req.PageIndex) * int(req.PageSize)).
		Limit(int(req.PageSize)).
		WithField().
		WithQuestions(func(qq *ent.QuestionQuery) {
			qq.Select(entquestion.FieldID, entquestion.FieldContent)
		}).
		All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	return scenarials, int32(totalCount), totalPage, nil
}

func (s *scenario) Get(ctx context.Context, req *ekko.GetScenarioRequest) (*ent.Scenario, error) {
	return s.ent.Scenario.Query().Where(entscenario.IDEQ(req.Id)).WithQuestions().WithField().First(ctx)
}

func (s *scenario) Favorite(ctx context.Context, userId uint64, scenarioId uint64) error {
	exist, err := s.ent.ScenarioFavorite.Query().Where(scenariofavorite.UserID(userId), scenariofavorite.ScenarioID(scenarioId)).Exist(ctx)
	if err != nil {
		return err
	}

	if exist {
		_, err := s.ent.ScenarioFavorite.Delete().Where(scenariofavorite.UserID(userId), scenariofavorite.ScenarioID(scenarioId)).Exec(ctx)
		return err
	}

	return s.ent.ScenarioFavorite.Create().
		SetUserID(userId).
		SetScenarioID(scenarioId).
		Exec(ctx)
}

func (s *scenario) Rating(ctx context.Context, id uint64, rating float64) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	return tx.WithTransaction(ctx, s.ent, func(ctx context.Context, tx tx.Tx) error {
		foundedScenario, err := tx.Client().Scenario.Query().Where(entscenario.IDEQ(id)).ForUpdate().Only(ctx)
		if err != nil {
			return err
		}

		newRating := (foundedScenario.Rating*float64(foundedScenario.TotalRating) + rating) / float64(foundedScenario.TotalRating+1)
		return foundedScenario.Update().SetRating(newRating).AddTotalRating(1).Exec(ctx)
	})
}
