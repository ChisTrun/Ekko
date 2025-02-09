// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ekko/package/ent/predicate"
	"ekko/package/ent/scenario"
	"ekko/package/ent/scenariocandidate"
	"ekko/package/ent/submissionattempt"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScenarioCandidateUpdate is the builder for updating ScenarioCandidate entities.
type ScenarioCandidateUpdate struct {
	config
	hooks     []Hook
	mutation  *ScenarioCandidateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ScenarioCandidateUpdate builder.
func (scu *ScenarioCandidateUpdate) Where(ps ...predicate.ScenarioCandidate) *ScenarioCandidateUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetUpdatedAt sets the "updated_at" field.
func (scu *ScenarioCandidateUpdate) SetUpdatedAt(t time.Time) *ScenarioCandidateUpdate {
	scu.mutation.SetUpdatedAt(t)
	return scu
}

// SetCandidateID sets the "candidate_id" field.
func (scu *ScenarioCandidateUpdate) SetCandidateID(u uint64) *ScenarioCandidateUpdate {
	scu.mutation.ResetCandidateID()
	scu.mutation.SetCandidateID(u)
	return scu
}

// SetNillableCandidateID sets the "candidate_id" field if the given value is not nil.
func (scu *ScenarioCandidateUpdate) SetNillableCandidateID(u *uint64) *ScenarioCandidateUpdate {
	if u != nil {
		scu.SetCandidateID(*u)
	}
	return scu
}

// AddCandidateID adds u to the "candidate_id" field.
func (scu *ScenarioCandidateUpdate) AddCandidateID(u int64) *ScenarioCandidateUpdate {
	scu.mutation.AddCandidateID(u)
	return scu
}

// SetScenarioID sets the "scenario_id" field.
func (scu *ScenarioCandidateUpdate) SetScenarioID(u uint64) *ScenarioCandidateUpdate {
	scu.mutation.SetScenarioID(u)
	return scu
}

// SetNillableScenarioID sets the "scenario_id" field if the given value is not nil.
func (scu *ScenarioCandidateUpdate) SetNillableScenarioID(u *uint64) *ScenarioCandidateUpdate {
	if u != nil {
		scu.SetScenarioID(*u)
	}
	return scu
}

// SetScenario sets the "scenario" edge to the Scenario entity.
func (scu *ScenarioCandidateUpdate) SetScenario(s *Scenario) *ScenarioCandidateUpdate {
	return scu.SetScenarioID(s.ID)
}

// AddAttemptIDs adds the "attempts" edge to the SubmissionAttempt entity by IDs.
func (scu *ScenarioCandidateUpdate) AddAttemptIDs(ids ...uint64) *ScenarioCandidateUpdate {
	scu.mutation.AddAttemptIDs(ids...)
	return scu
}

// AddAttempts adds the "attempts" edges to the SubmissionAttempt entity.
func (scu *ScenarioCandidateUpdate) AddAttempts(s ...*SubmissionAttempt) *ScenarioCandidateUpdate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scu.AddAttemptIDs(ids...)
}

// Mutation returns the ScenarioCandidateMutation object of the builder.
func (scu *ScenarioCandidateUpdate) Mutation() *ScenarioCandidateMutation {
	return scu.mutation
}

// ClearScenario clears the "scenario" edge to the Scenario entity.
func (scu *ScenarioCandidateUpdate) ClearScenario() *ScenarioCandidateUpdate {
	scu.mutation.ClearScenario()
	return scu
}

// ClearAttempts clears all "attempts" edges to the SubmissionAttempt entity.
func (scu *ScenarioCandidateUpdate) ClearAttempts() *ScenarioCandidateUpdate {
	scu.mutation.ClearAttempts()
	return scu
}

// RemoveAttemptIDs removes the "attempts" edge to SubmissionAttempt entities by IDs.
func (scu *ScenarioCandidateUpdate) RemoveAttemptIDs(ids ...uint64) *ScenarioCandidateUpdate {
	scu.mutation.RemoveAttemptIDs(ids...)
	return scu
}

// RemoveAttempts removes "attempts" edges to SubmissionAttempt entities.
func (scu *ScenarioCandidateUpdate) RemoveAttempts(s ...*SubmissionAttempt) *ScenarioCandidateUpdate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scu.RemoveAttemptIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *ScenarioCandidateUpdate) Save(ctx context.Context) (int, error) {
	scu.defaults()
	return withHooks(ctx, scu.sqlSave, scu.mutation, scu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scu *ScenarioCandidateUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *ScenarioCandidateUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *ScenarioCandidateUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scu *ScenarioCandidateUpdate) defaults() {
	if _, ok := scu.mutation.UpdatedAt(); !ok {
		v := scenariocandidate.UpdateDefaultUpdatedAt()
		scu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scu *ScenarioCandidateUpdate) check() error {
	if scu.mutation.ScenarioCleared() && len(scu.mutation.ScenarioIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ScenarioCandidate.scenario"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (scu *ScenarioCandidateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ScenarioCandidateUpdate {
	scu.modifiers = append(scu.modifiers, modifiers...)
	return scu
}

func (scu *ScenarioCandidateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := scu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(scenariocandidate.Table, scenariocandidate.Columns, sqlgraph.NewFieldSpec(scenariocandidate.FieldID, field.TypeUint64))
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.UpdatedAt(); ok {
		_spec.SetField(scenariocandidate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := scu.mutation.CandidateID(); ok {
		_spec.SetField(scenariocandidate.FieldCandidateID, field.TypeUint64, value)
	}
	if value, ok := scu.mutation.AddedCandidateID(); ok {
		_spec.AddField(scenariocandidate.FieldCandidateID, field.TypeUint64, value)
	}
	if scu.mutation.ScenarioCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scenariocandidate.ScenarioTable,
			Columns: []string{scenariocandidate.ScenarioColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scenario.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.ScenarioIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scenariocandidate.ScenarioTable,
			Columns: []string{scenariocandidate.ScenarioColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scenario.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scu.mutation.AttemptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenariocandidate.AttemptsTable,
			Columns: []string{scenariocandidate.AttemptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submissionattempt.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.RemovedAttemptsIDs(); len(nodes) > 0 && !scu.mutation.AttemptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenariocandidate.AttemptsTable,
			Columns: []string{scenariocandidate.AttemptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submissionattempt.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.AttemptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenariocandidate.AttemptsTable,
			Columns: []string{scenariocandidate.AttemptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submissionattempt.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(scu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scenariocandidate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	scu.mutation.done = true
	return n, nil
}

// ScenarioCandidateUpdateOne is the builder for updating a single ScenarioCandidate entity.
type ScenarioCandidateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ScenarioCandidateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (scuo *ScenarioCandidateUpdateOne) SetUpdatedAt(t time.Time) *ScenarioCandidateUpdateOne {
	scuo.mutation.SetUpdatedAt(t)
	return scuo
}

// SetCandidateID sets the "candidate_id" field.
func (scuo *ScenarioCandidateUpdateOne) SetCandidateID(u uint64) *ScenarioCandidateUpdateOne {
	scuo.mutation.ResetCandidateID()
	scuo.mutation.SetCandidateID(u)
	return scuo
}

// SetNillableCandidateID sets the "candidate_id" field if the given value is not nil.
func (scuo *ScenarioCandidateUpdateOne) SetNillableCandidateID(u *uint64) *ScenarioCandidateUpdateOne {
	if u != nil {
		scuo.SetCandidateID(*u)
	}
	return scuo
}

// AddCandidateID adds u to the "candidate_id" field.
func (scuo *ScenarioCandidateUpdateOne) AddCandidateID(u int64) *ScenarioCandidateUpdateOne {
	scuo.mutation.AddCandidateID(u)
	return scuo
}

// SetScenarioID sets the "scenario_id" field.
func (scuo *ScenarioCandidateUpdateOne) SetScenarioID(u uint64) *ScenarioCandidateUpdateOne {
	scuo.mutation.SetScenarioID(u)
	return scuo
}

// SetNillableScenarioID sets the "scenario_id" field if the given value is not nil.
func (scuo *ScenarioCandidateUpdateOne) SetNillableScenarioID(u *uint64) *ScenarioCandidateUpdateOne {
	if u != nil {
		scuo.SetScenarioID(*u)
	}
	return scuo
}

// SetScenario sets the "scenario" edge to the Scenario entity.
func (scuo *ScenarioCandidateUpdateOne) SetScenario(s *Scenario) *ScenarioCandidateUpdateOne {
	return scuo.SetScenarioID(s.ID)
}

// AddAttemptIDs adds the "attempts" edge to the SubmissionAttempt entity by IDs.
func (scuo *ScenarioCandidateUpdateOne) AddAttemptIDs(ids ...uint64) *ScenarioCandidateUpdateOne {
	scuo.mutation.AddAttemptIDs(ids...)
	return scuo
}

// AddAttempts adds the "attempts" edges to the SubmissionAttempt entity.
func (scuo *ScenarioCandidateUpdateOne) AddAttempts(s ...*SubmissionAttempt) *ScenarioCandidateUpdateOne {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scuo.AddAttemptIDs(ids...)
}

// Mutation returns the ScenarioCandidateMutation object of the builder.
func (scuo *ScenarioCandidateUpdateOne) Mutation() *ScenarioCandidateMutation {
	return scuo.mutation
}

// ClearScenario clears the "scenario" edge to the Scenario entity.
func (scuo *ScenarioCandidateUpdateOne) ClearScenario() *ScenarioCandidateUpdateOne {
	scuo.mutation.ClearScenario()
	return scuo
}

// ClearAttempts clears all "attempts" edges to the SubmissionAttempt entity.
func (scuo *ScenarioCandidateUpdateOne) ClearAttempts() *ScenarioCandidateUpdateOne {
	scuo.mutation.ClearAttempts()
	return scuo
}

// RemoveAttemptIDs removes the "attempts" edge to SubmissionAttempt entities by IDs.
func (scuo *ScenarioCandidateUpdateOne) RemoveAttemptIDs(ids ...uint64) *ScenarioCandidateUpdateOne {
	scuo.mutation.RemoveAttemptIDs(ids...)
	return scuo
}

// RemoveAttempts removes "attempts" edges to SubmissionAttempt entities.
func (scuo *ScenarioCandidateUpdateOne) RemoveAttempts(s ...*SubmissionAttempt) *ScenarioCandidateUpdateOne {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scuo.RemoveAttemptIDs(ids...)
}

// Where appends a list predicates to the ScenarioCandidateUpdate builder.
func (scuo *ScenarioCandidateUpdateOne) Where(ps ...predicate.ScenarioCandidate) *ScenarioCandidateUpdateOne {
	scuo.mutation.Where(ps...)
	return scuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *ScenarioCandidateUpdateOne) Select(field string, fields ...string) *ScenarioCandidateUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated ScenarioCandidate entity.
func (scuo *ScenarioCandidateUpdateOne) Save(ctx context.Context) (*ScenarioCandidate, error) {
	scuo.defaults()
	return withHooks(ctx, scuo.sqlSave, scuo.mutation, scuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *ScenarioCandidateUpdateOne) SaveX(ctx context.Context) *ScenarioCandidate {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *ScenarioCandidateUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *ScenarioCandidateUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scuo *ScenarioCandidateUpdateOne) defaults() {
	if _, ok := scuo.mutation.UpdatedAt(); !ok {
		v := scenariocandidate.UpdateDefaultUpdatedAt()
		scuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scuo *ScenarioCandidateUpdateOne) check() error {
	if scuo.mutation.ScenarioCleared() && len(scuo.mutation.ScenarioIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ScenarioCandidate.scenario"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (scuo *ScenarioCandidateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ScenarioCandidateUpdateOne {
	scuo.modifiers = append(scuo.modifiers, modifiers...)
	return scuo
}

func (scuo *ScenarioCandidateUpdateOne) sqlSave(ctx context.Context) (_node *ScenarioCandidate, err error) {
	if err := scuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(scenariocandidate.Table, scenariocandidate.Columns, sqlgraph.NewFieldSpec(scenariocandidate.FieldID, field.TypeUint64))
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ScenarioCandidate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scenariocandidate.FieldID)
		for _, f := range fields {
			if !scenariocandidate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != scenariocandidate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.UpdatedAt(); ok {
		_spec.SetField(scenariocandidate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := scuo.mutation.CandidateID(); ok {
		_spec.SetField(scenariocandidate.FieldCandidateID, field.TypeUint64, value)
	}
	if value, ok := scuo.mutation.AddedCandidateID(); ok {
		_spec.AddField(scenariocandidate.FieldCandidateID, field.TypeUint64, value)
	}
	if scuo.mutation.ScenarioCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scenariocandidate.ScenarioTable,
			Columns: []string{scenariocandidate.ScenarioColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scenario.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.ScenarioIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scenariocandidate.ScenarioTable,
			Columns: []string{scenariocandidate.ScenarioColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scenario.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scuo.mutation.AttemptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenariocandidate.AttemptsTable,
			Columns: []string{scenariocandidate.AttemptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submissionattempt.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.RemovedAttemptsIDs(); len(nodes) > 0 && !scuo.mutation.AttemptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenariocandidate.AttemptsTable,
			Columns: []string{scenariocandidate.AttemptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submissionattempt.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.AttemptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenariocandidate.AttemptsTable,
			Columns: []string{scenariocandidate.AttemptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submissionattempt.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(scuo.modifiers...)
	_node = &ScenarioCandidate{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scenariocandidate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	scuo.mutation.done = true
	return _node, nil
}
