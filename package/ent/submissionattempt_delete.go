// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ekko/package/ent/predicate"
	"ekko/package/ent/submissionattempt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubmissionAttemptDelete is the builder for deleting a SubmissionAttempt entity.
type SubmissionAttemptDelete struct {
	config
	hooks    []Hook
	mutation *SubmissionAttemptMutation
}

// Where appends a list predicates to the SubmissionAttemptDelete builder.
func (sad *SubmissionAttemptDelete) Where(ps ...predicate.SubmissionAttempt) *SubmissionAttemptDelete {
	sad.mutation.Where(ps...)
	return sad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sad *SubmissionAttemptDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sad.sqlExec, sad.mutation, sad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sad *SubmissionAttemptDelete) ExecX(ctx context.Context) int {
	n, err := sad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sad *SubmissionAttemptDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(submissionattempt.Table, sqlgraph.NewFieldSpec(submissionattempt.FieldID, field.TypeUint64))
	if ps := sad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sad.mutation.done = true
	return affected, err
}

// SubmissionAttemptDeleteOne is the builder for deleting a single SubmissionAttempt entity.
type SubmissionAttemptDeleteOne struct {
	sad *SubmissionAttemptDelete
}

// Where appends a list predicates to the SubmissionAttemptDelete builder.
func (sado *SubmissionAttemptDeleteOne) Where(ps ...predicate.SubmissionAttempt) *SubmissionAttemptDeleteOne {
	sado.sad.mutation.Where(ps...)
	return sado
}

// Exec executes the deletion query.
func (sado *SubmissionAttemptDeleteOne) Exec(ctx context.Context) error {
	n, err := sado.sad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{submissionattempt.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sado *SubmissionAttemptDeleteOne) ExecX(ctx context.Context) {
	if err := sado.Exec(ctx); err != nil {
		panic(err)
	}
}
