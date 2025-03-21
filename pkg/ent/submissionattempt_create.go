// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ekko/pkg/ent/answersubmission"
	"ekko/pkg/ent/scenariocandidate"
	"ekko/pkg/ent/submissionattempt"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubmissionAttemptCreate is the builder for creating a SubmissionAttempt entity.
type SubmissionAttemptCreate struct {
	config
	mutation *SubmissionAttemptMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sac *SubmissionAttemptCreate) SetCreatedAt(t time.Time) *SubmissionAttemptCreate {
	sac.mutation.SetCreatedAt(t)
	return sac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sac *SubmissionAttemptCreate) SetNillableCreatedAt(t *time.Time) *SubmissionAttemptCreate {
	if t != nil {
		sac.SetCreatedAt(*t)
	}
	return sac
}

// SetUpdatedAt sets the "updated_at" field.
func (sac *SubmissionAttemptCreate) SetUpdatedAt(t time.Time) *SubmissionAttemptCreate {
	sac.mutation.SetUpdatedAt(t)
	return sac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sac *SubmissionAttemptCreate) SetNillableUpdatedAt(t *time.Time) *SubmissionAttemptCreate {
	if t != nil {
		sac.SetUpdatedAt(*t)
	}
	return sac
}

// SetScenarioCandidateID sets the "scenario_candidate_id" field.
func (sac *SubmissionAttemptCreate) SetScenarioCandidateID(u uint64) *SubmissionAttemptCreate {
	sac.mutation.SetScenarioCandidateID(u)
	return sac
}

// SetAttemptNumber sets the "attempt_number" field.
func (sac *SubmissionAttemptCreate) SetAttemptNumber(i int32) *SubmissionAttemptCreate {
	sac.mutation.SetAttemptNumber(i)
	return sac
}

// SetID sets the "id" field.
func (sac *SubmissionAttemptCreate) SetID(u uint64) *SubmissionAttemptCreate {
	sac.mutation.SetID(u)
	return sac
}

// SetScenarioCandidate sets the "scenario_candidate" edge to the ScenarioCandidate entity.
func (sac *SubmissionAttemptCreate) SetScenarioCandidate(s *ScenarioCandidate) *SubmissionAttemptCreate {
	return sac.SetScenarioCandidateID(s.ID)
}

// AddAnswerIDs adds the "answers" edge to the AnswerSubmission entity by IDs.
func (sac *SubmissionAttemptCreate) AddAnswerIDs(ids ...uint64) *SubmissionAttemptCreate {
	sac.mutation.AddAnswerIDs(ids...)
	return sac
}

// AddAnswers adds the "answers" edges to the AnswerSubmission entity.
func (sac *SubmissionAttemptCreate) AddAnswers(a ...*AnswerSubmission) *SubmissionAttemptCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sac.AddAnswerIDs(ids...)
}

// Mutation returns the SubmissionAttemptMutation object of the builder.
func (sac *SubmissionAttemptCreate) Mutation() *SubmissionAttemptMutation {
	return sac.mutation
}

// Save creates the SubmissionAttempt in the database.
func (sac *SubmissionAttemptCreate) Save(ctx context.Context) (*SubmissionAttempt, error) {
	sac.defaults()
	return withHooks(ctx, sac.sqlSave, sac.mutation, sac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sac *SubmissionAttemptCreate) SaveX(ctx context.Context) *SubmissionAttempt {
	v, err := sac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sac *SubmissionAttemptCreate) Exec(ctx context.Context) error {
	_, err := sac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sac *SubmissionAttemptCreate) ExecX(ctx context.Context) {
	if err := sac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sac *SubmissionAttemptCreate) defaults() {
	if _, ok := sac.mutation.CreatedAt(); !ok {
		v := submissionattempt.DefaultCreatedAt()
		sac.mutation.SetCreatedAt(v)
	}
	if _, ok := sac.mutation.UpdatedAt(); !ok {
		v := submissionattempt.DefaultUpdatedAt()
		sac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sac *SubmissionAttemptCreate) check() error {
	if _, ok := sac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SubmissionAttempt.created_at"`)}
	}
	if _, ok := sac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SubmissionAttempt.updated_at"`)}
	}
	if _, ok := sac.mutation.ScenarioCandidateID(); !ok {
		return &ValidationError{Name: "scenario_candidate_id", err: errors.New(`ent: missing required field "SubmissionAttempt.scenario_candidate_id"`)}
	}
	if _, ok := sac.mutation.AttemptNumber(); !ok {
		return &ValidationError{Name: "attempt_number", err: errors.New(`ent: missing required field "SubmissionAttempt.attempt_number"`)}
	}
	if len(sac.mutation.ScenarioCandidateIDs()) == 0 {
		return &ValidationError{Name: "scenario_candidate", err: errors.New(`ent: missing required edge "SubmissionAttempt.scenario_candidate"`)}
	}
	return nil
}

func (sac *SubmissionAttemptCreate) sqlSave(ctx context.Context) (*SubmissionAttempt, error) {
	if err := sac.check(); err != nil {
		return nil, err
	}
	_node, _spec := sac.createSpec()
	if err := sqlgraph.CreateNode(ctx, sac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	sac.mutation.id = &_node.ID
	sac.mutation.done = true
	return _node, nil
}

func (sac *SubmissionAttemptCreate) createSpec() (*SubmissionAttempt, *sqlgraph.CreateSpec) {
	var (
		_node = &SubmissionAttempt{config: sac.config}
		_spec = sqlgraph.NewCreateSpec(submissionattempt.Table, sqlgraph.NewFieldSpec(submissionattempt.FieldID, field.TypeUint64))
	)
	_spec.OnConflict = sac.conflict
	if id, ok := sac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sac.mutation.CreatedAt(); ok {
		_spec.SetField(submissionattempt.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sac.mutation.UpdatedAt(); ok {
		_spec.SetField(submissionattempt.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sac.mutation.AttemptNumber(); ok {
		_spec.SetField(submissionattempt.FieldAttemptNumber, field.TypeInt32, value)
		_node.AttemptNumber = value
	}
	if nodes := sac.mutation.ScenarioCandidateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submissionattempt.ScenarioCandidateTable,
			Columns: []string{submissionattempt.ScenarioCandidateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scenariocandidate.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ScenarioCandidateID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sac.mutation.AnswersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submissionattempt.AnswersTable,
			Columns: []string{submissionattempt.AnswersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(answersubmission.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SubmissionAttempt.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubmissionAttemptUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (sac *SubmissionAttemptCreate) OnConflict(opts ...sql.ConflictOption) *SubmissionAttemptUpsertOne {
	sac.conflict = opts
	return &SubmissionAttemptUpsertOne{
		create: sac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SubmissionAttempt.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sac *SubmissionAttemptCreate) OnConflictColumns(columns ...string) *SubmissionAttemptUpsertOne {
	sac.conflict = append(sac.conflict, sql.ConflictColumns(columns...))
	return &SubmissionAttemptUpsertOne{
		create: sac,
	}
}

type (
	// SubmissionAttemptUpsertOne is the builder for "upsert"-ing
	//  one SubmissionAttempt node.
	SubmissionAttemptUpsertOne struct {
		create *SubmissionAttemptCreate
	}

	// SubmissionAttemptUpsert is the "OnConflict" setter.
	SubmissionAttemptUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *SubmissionAttemptUpsert) SetUpdatedAt(v time.Time) *SubmissionAttemptUpsert {
	u.Set(submissionattempt.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SubmissionAttemptUpsert) UpdateUpdatedAt() *SubmissionAttemptUpsert {
	u.SetExcluded(submissionattempt.FieldUpdatedAt)
	return u
}

// SetScenarioCandidateID sets the "scenario_candidate_id" field.
func (u *SubmissionAttemptUpsert) SetScenarioCandidateID(v uint64) *SubmissionAttemptUpsert {
	u.Set(submissionattempt.FieldScenarioCandidateID, v)
	return u
}

// UpdateScenarioCandidateID sets the "scenario_candidate_id" field to the value that was provided on create.
func (u *SubmissionAttemptUpsert) UpdateScenarioCandidateID() *SubmissionAttemptUpsert {
	u.SetExcluded(submissionattempt.FieldScenarioCandidateID)
	return u
}

// SetAttemptNumber sets the "attempt_number" field.
func (u *SubmissionAttemptUpsert) SetAttemptNumber(v int32) *SubmissionAttemptUpsert {
	u.Set(submissionattempt.FieldAttemptNumber, v)
	return u
}

// UpdateAttemptNumber sets the "attempt_number" field to the value that was provided on create.
func (u *SubmissionAttemptUpsert) UpdateAttemptNumber() *SubmissionAttemptUpsert {
	u.SetExcluded(submissionattempt.FieldAttemptNumber)
	return u
}

// AddAttemptNumber adds v to the "attempt_number" field.
func (u *SubmissionAttemptUpsert) AddAttemptNumber(v int32) *SubmissionAttemptUpsert {
	u.Add(submissionattempt.FieldAttemptNumber, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SubmissionAttempt.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(submissionattempt.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubmissionAttemptUpsertOne) UpdateNewValues() *SubmissionAttemptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(submissionattempt.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(submissionattempt.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SubmissionAttempt.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SubmissionAttemptUpsertOne) Ignore() *SubmissionAttemptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubmissionAttemptUpsertOne) DoNothing() *SubmissionAttemptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubmissionAttemptCreate.OnConflict
// documentation for more info.
func (u *SubmissionAttemptUpsertOne) Update(set func(*SubmissionAttemptUpsert)) *SubmissionAttemptUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubmissionAttemptUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SubmissionAttemptUpsertOne) SetUpdatedAt(v time.Time) *SubmissionAttemptUpsertOne {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SubmissionAttemptUpsertOne) UpdateUpdatedAt() *SubmissionAttemptUpsertOne {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetScenarioCandidateID sets the "scenario_candidate_id" field.
func (u *SubmissionAttemptUpsertOne) SetScenarioCandidateID(v uint64) *SubmissionAttemptUpsertOne {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.SetScenarioCandidateID(v)
	})
}

// UpdateScenarioCandidateID sets the "scenario_candidate_id" field to the value that was provided on create.
func (u *SubmissionAttemptUpsertOne) UpdateScenarioCandidateID() *SubmissionAttemptUpsertOne {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.UpdateScenarioCandidateID()
	})
}

// SetAttemptNumber sets the "attempt_number" field.
func (u *SubmissionAttemptUpsertOne) SetAttemptNumber(v int32) *SubmissionAttemptUpsertOne {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.SetAttemptNumber(v)
	})
}

// AddAttemptNumber adds v to the "attempt_number" field.
func (u *SubmissionAttemptUpsertOne) AddAttemptNumber(v int32) *SubmissionAttemptUpsertOne {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.AddAttemptNumber(v)
	})
}

// UpdateAttemptNumber sets the "attempt_number" field to the value that was provided on create.
func (u *SubmissionAttemptUpsertOne) UpdateAttemptNumber() *SubmissionAttemptUpsertOne {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.UpdateAttemptNumber()
	})
}

// Exec executes the query.
func (u *SubmissionAttemptUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SubmissionAttemptCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubmissionAttemptUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SubmissionAttemptUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SubmissionAttemptUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SubmissionAttemptCreateBulk is the builder for creating many SubmissionAttempt entities in bulk.
type SubmissionAttemptCreateBulk struct {
	config
	err      error
	builders []*SubmissionAttemptCreate
	conflict []sql.ConflictOption
}

// Save creates the SubmissionAttempt entities in the database.
func (sacb *SubmissionAttemptCreateBulk) Save(ctx context.Context) ([]*SubmissionAttempt, error) {
	if sacb.err != nil {
		return nil, sacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sacb.builders))
	nodes := make([]*SubmissionAttempt, len(sacb.builders))
	mutators := make([]Mutator, len(sacb.builders))
	for i := range sacb.builders {
		func(i int, root context.Context) {
			builder := sacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubmissionAttemptMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sacb *SubmissionAttemptCreateBulk) SaveX(ctx context.Context) []*SubmissionAttempt {
	v, err := sacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sacb *SubmissionAttemptCreateBulk) Exec(ctx context.Context) error {
	_, err := sacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sacb *SubmissionAttemptCreateBulk) ExecX(ctx context.Context) {
	if err := sacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SubmissionAttempt.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubmissionAttemptUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (sacb *SubmissionAttemptCreateBulk) OnConflict(opts ...sql.ConflictOption) *SubmissionAttemptUpsertBulk {
	sacb.conflict = opts
	return &SubmissionAttemptUpsertBulk{
		create: sacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SubmissionAttempt.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sacb *SubmissionAttemptCreateBulk) OnConflictColumns(columns ...string) *SubmissionAttemptUpsertBulk {
	sacb.conflict = append(sacb.conflict, sql.ConflictColumns(columns...))
	return &SubmissionAttemptUpsertBulk{
		create: sacb,
	}
}

// SubmissionAttemptUpsertBulk is the builder for "upsert"-ing
// a bulk of SubmissionAttempt nodes.
type SubmissionAttemptUpsertBulk struct {
	create *SubmissionAttemptCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SubmissionAttempt.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(submissionattempt.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubmissionAttemptUpsertBulk) UpdateNewValues() *SubmissionAttemptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(submissionattempt.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(submissionattempt.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SubmissionAttempt.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SubmissionAttemptUpsertBulk) Ignore() *SubmissionAttemptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubmissionAttemptUpsertBulk) DoNothing() *SubmissionAttemptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubmissionAttemptCreateBulk.OnConflict
// documentation for more info.
func (u *SubmissionAttemptUpsertBulk) Update(set func(*SubmissionAttemptUpsert)) *SubmissionAttemptUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubmissionAttemptUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SubmissionAttemptUpsertBulk) SetUpdatedAt(v time.Time) *SubmissionAttemptUpsertBulk {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SubmissionAttemptUpsertBulk) UpdateUpdatedAt() *SubmissionAttemptUpsertBulk {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetScenarioCandidateID sets the "scenario_candidate_id" field.
func (u *SubmissionAttemptUpsertBulk) SetScenarioCandidateID(v uint64) *SubmissionAttemptUpsertBulk {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.SetScenarioCandidateID(v)
	})
}

// UpdateScenarioCandidateID sets the "scenario_candidate_id" field to the value that was provided on create.
func (u *SubmissionAttemptUpsertBulk) UpdateScenarioCandidateID() *SubmissionAttemptUpsertBulk {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.UpdateScenarioCandidateID()
	})
}

// SetAttemptNumber sets the "attempt_number" field.
func (u *SubmissionAttemptUpsertBulk) SetAttemptNumber(v int32) *SubmissionAttemptUpsertBulk {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.SetAttemptNumber(v)
	})
}

// AddAttemptNumber adds v to the "attempt_number" field.
func (u *SubmissionAttemptUpsertBulk) AddAttemptNumber(v int32) *SubmissionAttemptUpsertBulk {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.AddAttemptNumber(v)
	})
}

// UpdateAttemptNumber sets the "attempt_number" field to the value that was provided on create.
func (u *SubmissionAttemptUpsertBulk) UpdateAttemptNumber() *SubmissionAttemptUpsertBulk {
	return u.Update(func(s *SubmissionAttemptUpsert) {
		s.UpdateAttemptNumber()
	})
}

// Exec executes the query.
func (u *SubmissionAttemptUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SubmissionAttemptCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SubmissionAttemptCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubmissionAttemptUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
