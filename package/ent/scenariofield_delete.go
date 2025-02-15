// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ekko/package/ent/predicate"
	"ekko/package/ent/scenariofield"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScenarioFieldDelete is the builder for deleting a ScenarioField entity.
type ScenarioFieldDelete struct {
	config
	hooks    []Hook
	mutation *ScenarioFieldMutation
}

// Where appends a list predicates to the ScenarioFieldDelete builder.
func (sfd *ScenarioFieldDelete) Where(ps ...predicate.ScenarioField) *ScenarioFieldDelete {
	sfd.mutation.Where(ps...)
	return sfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sfd *ScenarioFieldDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sfd.sqlExec, sfd.mutation, sfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sfd *ScenarioFieldDelete) ExecX(ctx context.Context) int {
	n, err := sfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sfd *ScenarioFieldDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(scenariofield.Table, sqlgraph.NewFieldSpec(scenariofield.FieldID, field.TypeUint64))
	if ps := sfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sfd.mutation.done = true
	return affected, err
}

// ScenarioFieldDeleteOne is the builder for deleting a single ScenarioField entity.
type ScenarioFieldDeleteOne struct {
	sfd *ScenarioFieldDelete
}

// Where appends a list predicates to the ScenarioFieldDelete builder.
func (sfdo *ScenarioFieldDeleteOne) Where(ps ...predicate.ScenarioField) *ScenarioFieldDeleteOne {
	sfdo.sfd.mutation.Where(ps...)
	return sfdo
}

// Exec executes the deletion query.
func (sfdo *ScenarioFieldDeleteOne) Exec(ctx context.Context) error {
	n, err := sfdo.sfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{scenariofield.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sfdo *ScenarioFieldDeleteOne) ExecX(ctx context.Context) {
	if err := sfdo.Exec(ctx); err != nil {
		panic(err)
	}
}
