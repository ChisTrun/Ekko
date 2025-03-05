package field

import (
	"context"
	ekko "ekko/api"
	"ekko/internal/utils/paging"
	utils "ekko/internal/utils/sort"
	"ekko/internal/utils/tx"
	"ekko/pkg/ent"
	"ekko/pkg/ent/scenariofield"
)

type Field interface {
	Create(ctx context.Context, req *ekko.CreateFieldRequest) (*ent.ScenarioField, error)
	Update(ctx context.Context, tx tx.Tx, req *ekko.UpdateFieldRequest) error
	Delete(ctx context.Context, tx tx.Tx, req *ekko.DeleteFieldRequest) error
	List(ctx context.Context, req *ekko.ListFieldRequest) ([]*ent.ScenarioField, int32, int32, error)
}

type field struct {
	ent *ent.Client
}

func New(ent *ent.Client) Field {
	return &field{
		ent: ent,
	}
}

func (f *field) Create(ctx context.Context, req *ekko.CreateFieldRequest) (*ent.ScenarioField, error) {
	return f.ent.ScenarioField.Create().SetName(req.Name).Save(ctx)
}

func (f *field) Update(ctx context.Context, tx tx.Tx, req *ekko.UpdateFieldRequest) error {
	_, err := tx.Client().ScenarioField.UpdateOneID(req.Id).SetName(req.Name).Save(ctx)
	return err
}

func (f *field) Delete(ctx context.Context, tx tx.Tx, req *ekko.DeleteFieldRequest) error {
	_, err := tx.Client().ScenarioField.Delete().Where(scenariofield.IDIn(req.Ids...)).Exec(ctx)
	return err
}

func (f *field) List(ctx context.Context, req *ekko.ListFieldRequest) ([]*ent.ScenarioField, int32, int32, error) {
	query := f.ent.ScenarioField.Query()

	if len(req.Ids) > 0 {
		query = query.Where(scenariofield.IDIn(req.Ids...))
	}

	if req.SearchContent != nil {
		query = query.Where(scenariofield.NameContainsFold(*req.SearchContent))
	}

	sort, err := utils.GetSort(scenariofield.Columns, scenariofield.Table, req.GetSortMethods())
	if err != nil {
		return nil, 0, 0, err
	}
	totalCount, _ := query.Count(ctx)
	totalPage := paging.GetPagingData(int32(totalCount), req.PageSize)

	fields, err := query.
		Modify(sort).
		Offset(int(req.PageIndex) * int(req.PageSize)).
		Limit(int(req.PageSize)).
		All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	return fields, int32(totalCount), totalPage, nil

}
