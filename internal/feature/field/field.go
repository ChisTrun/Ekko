package field

import (
	"context"
	ekko "ekko/api"
	"ekko/internal/repository"
	"ekko/internal/utils/checker"
	"ekko/internal/utils/converter"
	"ekko/internal/utils/extractor"
	"ekko/internal/utils/tx"
	"ekko/package/ent"
	bulbasaur "ekko/third_party/bulbasaur/api"
	"sync"
)

type Field interface {
	CreateField(ctx context.Context, request *ekko.CreateFieldRequest) (*ekko.CreateFieldResponse, error)
	UpdateField(ctx context.Context, request *ekko.UpdateFieldRequest) error
	DeleteField(ctx context.Context, request *ekko.DeleteFieldRequest) error
	ListField(ctx context.Context, request *ekko.ListFieldRequest) (*ekko.ListFieldResponse, error)
}

type field struct {
	repo      *repository.Repository
	extractor extractor.Extractor
}

func New(repo *repository.Repository, extractor extractor.Extractor) Field {
	return &field{
		repo:      repo,
		extractor: extractor,
	}
}

func (f *field) CreateField(ctx context.Context, request *ekko.CreateFieldRequest) (*ekko.CreateFieldResponse, error) {
	roleIds := f.extractor.GetRoleIDs(ctx)
	if err := checker.CheckRole(ctx, int32(bulbasaur.Role_ROLE_BUSINESS_MANAGER), roleIds); err != nil {
		return nil, err
	}

	entField, err := f.repo.Field.Create(ctx, request)
	if err != nil {
		return nil, err
	}

	return &ekko.CreateFieldResponse{
		Field: converter.ConvertField(entField),
	}, nil
}

func (f *field) UpdateField(ctx context.Context, request *ekko.UpdateFieldRequest) error {
	roleIds := f.extractor.GetRoleIDs(ctx)
	if err := checker.CheckRole(ctx, int32(bulbasaur.Role_ROLE_BUSINESS_MANAGER), roleIds); err != nil {
		return err
	}

	return tx.WithTransaction(ctx, f.repo.Ent, func(ctx context.Context, tx tx.Tx) error {
		return f.repo.Field.Update(ctx, tx, request)
	})
}

func (f *field) DeleteField(ctx context.Context, request *ekko.DeleteFieldRequest) error {
	roleIds := f.extractor.GetRoleIDs(ctx)
	if err := checker.CheckRole(ctx, int32(bulbasaur.Role_ROLE_BUSINESS_MANAGER), roleIds); err != nil {
		return err
	}

	return tx.WithTransaction(ctx, f.repo.Ent, func(ctx context.Context, tx tx.Tx) error {
		return f.repo.Field.Delete(ctx, tx, request)
	})
}

func (f *field) ListField(ctx context.Context, request *ekko.ListFieldRequest) (*ekko.ListFieldResponse, error) {
	fields, totalCount, totalPage, err := f.repo.Field.List(ctx, request)
	if err != nil {
		return nil, err
	}

	data := make([]*ekko.Field, len(fields))
	var wg sync.WaitGroup

	for i, field := range fields {
		wg.Add(1)
		go func(i int, field ent.ScenarioField) {
			defer wg.Done()
			data[i] = converter.ConvertField(&field)
		}(i, *field)
	}
	wg.Wait()

	return &ekko.ListFieldResponse{
		Fields:     data,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}, nil
}
