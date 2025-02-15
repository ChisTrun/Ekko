// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ekko/package/ent/scenario"
	"ekko/package/ent/scenariofield"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScenarioFieldCreate is the builder for creating a ScenarioField entity.
type ScenarioFieldCreate struct {
	config
	mutation *ScenarioFieldMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sfc *ScenarioFieldCreate) SetCreatedAt(t time.Time) *ScenarioFieldCreate {
	sfc.mutation.SetCreatedAt(t)
	return sfc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sfc *ScenarioFieldCreate) SetNillableCreatedAt(t *time.Time) *ScenarioFieldCreate {
	if t != nil {
		sfc.SetCreatedAt(*t)
	}
	return sfc
}

// SetUpdatedAt sets the "updated_at" field.
func (sfc *ScenarioFieldCreate) SetUpdatedAt(t time.Time) *ScenarioFieldCreate {
	sfc.mutation.SetUpdatedAt(t)
	return sfc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sfc *ScenarioFieldCreate) SetNillableUpdatedAt(t *time.Time) *ScenarioFieldCreate {
	if t != nil {
		sfc.SetUpdatedAt(*t)
	}
	return sfc
}

// SetName sets the "name" field.
func (sfc *ScenarioFieldCreate) SetName(s string) *ScenarioFieldCreate {
	sfc.mutation.SetName(s)
	return sfc
}

// SetID sets the "id" field.
func (sfc *ScenarioFieldCreate) SetID(u uint64) *ScenarioFieldCreate {
	sfc.mutation.SetID(u)
	return sfc
}

// SetSenariosID sets the "senarios" edge to the Scenario entity by ID.
func (sfc *ScenarioFieldCreate) SetSenariosID(id uint64) *ScenarioFieldCreate {
	sfc.mutation.SetSenariosID(id)
	return sfc
}

// SetNillableSenariosID sets the "senarios" edge to the Scenario entity by ID if the given value is not nil.
func (sfc *ScenarioFieldCreate) SetNillableSenariosID(id *uint64) *ScenarioFieldCreate {
	if id != nil {
		sfc = sfc.SetSenariosID(*id)
	}
	return sfc
}

// SetSenarios sets the "senarios" edge to the Scenario entity.
func (sfc *ScenarioFieldCreate) SetSenarios(s *Scenario) *ScenarioFieldCreate {
	return sfc.SetSenariosID(s.ID)
}

// Mutation returns the ScenarioFieldMutation object of the builder.
func (sfc *ScenarioFieldCreate) Mutation() *ScenarioFieldMutation {
	return sfc.mutation
}

// Save creates the ScenarioField in the database.
func (sfc *ScenarioFieldCreate) Save(ctx context.Context) (*ScenarioField, error) {
	sfc.defaults()
	return withHooks(ctx, sfc.sqlSave, sfc.mutation, sfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sfc *ScenarioFieldCreate) SaveX(ctx context.Context) *ScenarioField {
	v, err := sfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sfc *ScenarioFieldCreate) Exec(ctx context.Context) error {
	_, err := sfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sfc *ScenarioFieldCreate) ExecX(ctx context.Context) {
	if err := sfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sfc *ScenarioFieldCreate) defaults() {
	if _, ok := sfc.mutation.CreatedAt(); !ok {
		v := scenariofield.DefaultCreatedAt()
		sfc.mutation.SetCreatedAt(v)
	}
	if _, ok := sfc.mutation.UpdatedAt(); !ok {
		v := scenariofield.DefaultUpdatedAt()
		sfc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sfc *ScenarioFieldCreate) check() error {
	if _, ok := sfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ScenarioField.created_at"`)}
	}
	if _, ok := sfc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ScenarioField.updated_at"`)}
	}
	if _, ok := sfc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ScenarioField.name"`)}
	}
	return nil
}

func (sfc *ScenarioFieldCreate) sqlSave(ctx context.Context) (*ScenarioField, error) {
	if err := sfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	sfc.mutation.id = &_node.ID
	sfc.mutation.done = true
	return _node, nil
}

func (sfc *ScenarioFieldCreate) createSpec() (*ScenarioField, *sqlgraph.CreateSpec) {
	var (
		_node = &ScenarioField{config: sfc.config}
		_spec = sqlgraph.NewCreateSpec(scenariofield.Table, sqlgraph.NewFieldSpec(scenariofield.FieldID, field.TypeUint64))
	)
	_spec.OnConflict = sfc.conflict
	if id, ok := sfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sfc.mutation.CreatedAt(); ok {
		_spec.SetField(scenariofield.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sfc.mutation.UpdatedAt(); ok {
		_spec.SetField(scenariofield.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sfc.mutation.Name(); ok {
		_spec.SetField(scenariofield.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := sfc.mutation.SenariosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   scenariofield.SenariosTable,
			Columns: []string{scenariofield.SenariosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scenario.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.scenario_field_senarios = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ScenarioField.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ScenarioFieldUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (sfc *ScenarioFieldCreate) OnConflict(opts ...sql.ConflictOption) *ScenarioFieldUpsertOne {
	sfc.conflict = opts
	return &ScenarioFieldUpsertOne{
		create: sfc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ScenarioField.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sfc *ScenarioFieldCreate) OnConflictColumns(columns ...string) *ScenarioFieldUpsertOne {
	sfc.conflict = append(sfc.conflict, sql.ConflictColumns(columns...))
	return &ScenarioFieldUpsertOne{
		create: sfc,
	}
}

type (
	// ScenarioFieldUpsertOne is the builder for "upsert"-ing
	//  one ScenarioField node.
	ScenarioFieldUpsertOne struct {
		create *ScenarioFieldCreate
	}

	// ScenarioFieldUpsert is the "OnConflict" setter.
	ScenarioFieldUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ScenarioFieldUpsert) SetUpdatedAt(v time.Time) *ScenarioFieldUpsert {
	u.Set(scenariofield.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ScenarioFieldUpsert) UpdateUpdatedAt() *ScenarioFieldUpsert {
	u.SetExcluded(scenariofield.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *ScenarioFieldUpsert) SetName(v string) *ScenarioFieldUpsert {
	u.Set(scenariofield.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ScenarioFieldUpsert) UpdateName() *ScenarioFieldUpsert {
	u.SetExcluded(scenariofield.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ScenarioField.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(scenariofield.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ScenarioFieldUpsertOne) UpdateNewValues() *ScenarioFieldUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(scenariofield.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(scenariofield.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ScenarioField.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ScenarioFieldUpsertOne) Ignore() *ScenarioFieldUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ScenarioFieldUpsertOne) DoNothing() *ScenarioFieldUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ScenarioFieldCreate.OnConflict
// documentation for more info.
func (u *ScenarioFieldUpsertOne) Update(set func(*ScenarioFieldUpsert)) *ScenarioFieldUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ScenarioFieldUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ScenarioFieldUpsertOne) SetUpdatedAt(v time.Time) *ScenarioFieldUpsertOne {
	return u.Update(func(s *ScenarioFieldUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ScenarioFieldUpsertOne) UpdateUpdatedAt() *ScenarioFieldUpsertOne {
	return u.Update(func(s *ScenarioFieldUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *ScenarioFieldUpsertOne) SetName(v string) *ScenarioFieldUpsertOne {
	return u.Update(func(s *ScenarioFieldUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ScenarioFieldUpsertOne) UpdateName() *ScenarioFieldUpsertOne {
	return u.Update(func(s *ScenarioFieldUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *ScenarioFieldUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ScenarioFieldCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ScenarioFieldUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ScenarioFieldUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ScenarioFieldUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ScenarioFieldCreateBulk is the builder for creating many ScenarioField entities in bulk.
type ScenarioFieldCreateBulk struct {
	config
	err      error
	builders []*ScenarioFieldCreate
	conflict []sql.ConflictOption
}

// Save creates the ScenarioField entities in the database.
func (sfcb *ScenarioFieldCreateBulk) Save(ctx context.Context) ([]*ScenarioField, error) {
	if sfcb.err != nil {
		return nil, sfcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sfcb.builders))
	nodes := make([]*ScenarioField, len(sfcb.builders))
	mutators := make([]Mutator, len(sfcb.builders))
	for i := range sfcb.builders {
		func(i int, root context.Context) {
			builder := sfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScenarioFieldMutation)
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
					_, err = mutators[i+1].Mutate(root, sfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sfcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sfcb *ScenarioFieldCreateBulk) SaveX(ctx context.Context) []*ScenarioField {
	v, err := sfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sfcb *ScenarioFieldCreateBulk) Exec(ctx context.Context) error {
	_, err := sfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sfcb *ScenarioFieldCreateBulk) ExecX(ctx context.Context) {
	if err := sfcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ScenarioField.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ScenarioFieldUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (sfcb *ScenarioFieldCreateBulk) OnConflict(opts ...sql.ConflictOption) *ScenarioFieldUpsertBulk {
	sfcb.conflict = opts
	return &ScenarioFieldUpsertBulk{
		create: sfcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ScenarioField.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sfcb *ScenarioFieldCreateBulk) OnConflictColumns(columns ...string) *ScenarioFieldUpsertBulk {
	sfcb.conflict = append(sfcb.conflict, sql.ConflictColumns(columns...))
	return &ScenarioFieldUpsertBulk{
		create: sfcb,
	}
}

// ScenarioFieldUpsertBulk is the builder for "upsert"-ing
// a bulk of ScenarioField nodes.
type ScenarioFieldUpsertBulk struct {
	create *ScenarioFieldCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ScenarioField.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(scenariofield.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ScenarioFieldUpsertBulk) UpdateNewValues() *ScenarioFieldUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(scenariofield.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(scenariofield.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ScenarioField.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ScenarioFieldUpsertBulk) Ignore() *ScenarioFieldUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ScenarioFieldUpsertBulk) DoNothing() *ScenarioFieldUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ScenarioFieldCreateBulk.OnConflict
// documentation for more info.
func (u *ScenarioFieldUpsertBulk) Update(set func(*ScenarioFieldUpsert)) *ScenarioFieldUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ScenarioFieldUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ScenarioFieldUpsertBulk) SetUpdatedAt(v time.Time) *ScenarioFieldUpsertBulk {
	return u.Update(func(s *ScenarioFieldUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ScenarioFieldUpsertBulk) UpdateUpdatedAt() *ScenarioFieldUpsertBulk {
	return u.Update(func(s *ScenarioFieldUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *ScenarioFieldUpsertBulk) SetName(v string) *ScenarioFieldUpsertBulk {
	return u.Update(func(s *ScenarioFieldUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ScenarioFieldUpsertBulk) UpdateName() *ScenarioFieldUpsertBulk {
	return u.Update(func(s *ScenarioFieldUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *ScenarioFieldUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ScenarioFieldCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ScenarioFieldCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ScenarioFieldUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
