// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ekko/pkg/ent/answersubmission"
	"ekko/pkg/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AnswerSubmissionDelete is the builder for deleting a AnswerSubmission entity.
type AnswerSubmissionDelete struct {
	config
	hooks    []Hook
	mutation *AnswerSubmissionMutation
}

// Where appends a list predicates to the AnswerSubmissionDelete builder.
func (asd *AnswerSubmissionDelete) Where(ps ...predicate.AnswerSubmission) *AnswerSubmissionDelete {
	asd.mutation.Where(ps...)
	return asd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (asd *AnswerSubmissionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, asd.sqlExec, asd.mutation, asd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (asd *AnswerSubmissionDelete) ExecX(ctx context.Context) int {
	n, err := asd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (asd *AnswerSubmissionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(answersubmission.Table, sqlgraph.NewFieldSpec(answersubmission.FieldID, field.TypeUint64))
	if ps := asd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, asd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	asd.mutation.done = true
	return affected, err
}

// AnswerSubmissionDeleteOne is the builder for deleting a single AnswerSubmission entity.
type AnswerSubmissionDeleteOne struct {
	asd *AnswerSubmissionDelete
}

// Where appends a list predicates to the AnswerSubmissionDelete builder.
func (asdo *AnswerSubmissionDeleteOne) Where(ps ...predicate.AnswerSubmission) *AnswerSubmissionDeleteOne {
	asdo.asd.mutation.Where(ps...)
	return asdo
}

// Exec executes the deletion query.
func (asdo *AnswerSubmissionDeleteOne) Exec(ctx context.Context) error {
	n, err := asdo.asd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{answersubmission.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (asdo *AnswerSubmissionDeleteOne) ExecX(ctx context.Context) {
	if err := asdo.Exec(ctx); err != nil {
		panic(err)
	}
}
