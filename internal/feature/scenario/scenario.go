package scenario

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
	"fmt"
	"sync"
)

type Scenario interface {
	CreateScenario(ctx context.Context, request *ekko.CreateScenarioRequest) (*ekko.CreateScenarioResponse, error)
	UpdateScenario(ctx context.Context, request *ekko.UpdateScenarioRequest) error
	DeleteScenario(ctx context.Context, request *ekko.DeleteScenarioRequest) error
	ListScenario(ctx context.Context, isAuth bool, req *ekko.ListScenarioRequest) (*ekko.ListScenarioResponse, error)
	GetScenario(ctx context.Context, req *ekko.GetScenarioRequest) (*ekko.GetScenarioResponse, error)
	FavoriteScenario(ctx context.Context, req *ekko.FavoriteScenarioRequest) error
	RatingScenario(ctx context.Context, req *ekko.RatingScenarioRequest) error
}

type scenario struct {
	repo      *repository.Repository
	extractor extractor.Extractor
}

func New(repo *repository.Repository, extractor extractor.Extractor) Scenario {
	return &scenario{
		repo:      repo,
		extractor: extractor,
	}
}

func (s *scenario) CreateScenario(ctx context.Context, request *ekko.CreateScenarioRequest) (*ekko.CreateScenarioResponse, error) {
	roleIds := s.extractor.GetRoleIDs(ctx)
	if err := checker.CheckRole(ctx, fmt.Sprintf("%v", int32(bulbasaur.Role_ROLE_BUSINESS_MANAGER)), roleIds); err != nil {
		return nil, err
	}

	bmId, err := s.extractor.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	var entScenario *ent.Scenario
	if txErr := tx.WithTransaction(ctx, s.repo.Ent, func(ctx context.Context, tx tx.Tx) error {
		var err error
		entScenario, err = s.repo.Scenario.Create(ctx, tx, uint64(bmId), request)
		return err
	}); txErr != nil {
		return nil, txErr
	}

	return &ekko.CreateScenarioResponse{
		Scenario: converter.ConvertScenario(entScenario),
	}, nil

}

func (s *scenario) UpdateScenario(ctx context.Context, request *ekko.UpdateScenarioRequest) error {
	roleIds := s.extractor.GetRoleIDs(ctx)
	if err := checker.CheckRole(ctx, fmt.Sprintf("%v", int32(bulbasaur.Role_ROLE_BUSINESS_MANAGER)), roleIds); err != nil {
		return err
	}

	bmId, err := s.extractor.GetUserID(ctx)
	if err != nil {
		return err
	}

	return tx.WithTransaction(ctx, s.repo.Ent, func(ctx context.Context, tx tx.Tx) error {
		return s.repo.Scenario.Update(ctx, tx, uint64(bmId), request)
	})
}

func (s *scenario) DeleteScenario(ctx context.Context, request *ekko.DeleteScenarioRequest) error {
	roleIds := s.extractor.GetRoleIDs(ctx)
	if err := checker.CheckRole(ctx, fmt.Sprintf("%v", int32(bulbasaur.Role_ROLE_BUSINESS_MANAGER)), roleIds); err != nil {
		return err
	}

	bmId, err := s.extractor.GetUserID(ctx)
	if err != nil {
		return err
	}

	return tx.WithTransaction(ctx, s.repo.Ent, func(ctx context.Context, tx tx.Tx) error {
		return s.repo.Scenario.Delete(ctx, tx, uint64(bmId), request.Ids)
	})
}

func (s *scenario) ListScenario(ctx context.Context, isAuth bool, req *ekko.ListScenarioRequest) (*ekko.ListScenarioResponse, error) {

	var convertedUserId *uint64 = nil
	if isAuth {
		userId, err := s.extractor.GetUserID(ctx)
		if err != nil {
			return nil, err
		}
		tmp := uint64(userId)
		convertedUserId = &tmp
	}

	entScenarios, totalCount, totalPage, err := s.repo.Scenario.List(ctx, req, convertedUserId)
	if err != nil {
		return nil, err
	}

	data := make([]*ekko.Scenario, len(entScenarios))
	var wg sync.WaitGroup

	for i, scenario := range entScenarios {
		wg.Add(1)
		go func(i int, scenario ent.Scenario) {
			defer wg.Done()
			data[i] = converter.ConvertScenario(&scenario)
		}(i, *scenario)
	}
	wg.Wait()

	return &ekko.ListScenarioResponse{
		Scenario:   data,
		TotalCount: totalCount,
		TotalPage:  totalPage,
		Request:    req,
	}, nil
}

func (s *scenario) GetScenario(ctx context.Context, req *ekko.GetScenarioRequest) (*ekko.GetScenarioResponse, error) {

	entScenario, err := s.repo.Scenario.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return &ekko.GetScenarioResponse{
		Scenario: converter.ConvertScenario(entScenario),
	}, nil
}

func (s *scenario) FavoriteScenario(ctx context.Context, req *ekko.FavoriteScenarioRequest) error {

	userId, err := s.extractor.GetUserID(ctx)
	if err != nil {
		return err
	}

	return s.repo.Scenario.Favorite(ctx, uint64(userId), req.GetId())
}

func (s *scenario) RatingScenario(ctx context.Context, req *ekko.RatingScenarioRequest) error {
	return s.repo.Scenario.Rating(ctx, req.GetId(), float64(req.GetRating()))
}
