// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"ekko/pkg/ent/predicate"
	"ekko/pkg/ent/scenario"
	"ekko/pkg/ent/scenariocandidate"
	"ekko/pkg/ent/submissionattempt"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScenarioCandidateQuery is the builder for querying ScenarioCandidate entities.
type ScenarioCandidateQuery struct {
	config
	ctx          *QueryContext
	order        []scenariocandidate.OrderOption
	inters       []Interceptor
	predicates   []predicate.ScenarioCandidate
	withScenario *ScenarioQuery
	withAttempts *SubmissionAttemptQuery
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ScenarioCandidateQuery builder.
func (scq *ScenarioCandidateQuery) Where(ps ...predicate.ScenarioCandidate) *ScenarioCandidateQuery {
	scq.predicates = append(scq.predicates, ps...)
	return scq
}

// Limit the number of records to be returned by this query.
func (scq *ScenarioCandidateQuery) Limit(limit int) *ScenarioCandidateQuery {
	scq.ctx.Limit = &limit
	return scq
}

// Offset to start from.
func (scq *ScenarioCandidateQuery) Offset(offset int) *ScenarioCandidateQuery {
	scq.ctx.Offset = &offset
	return scq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (scq *ScenarioCandidateQuery) Unique(unique bool) *ScenarioCandidateQuery {
	scq.ctx.Unique = &unique
	return scq
}

// Order specifies how the records should be ordered.
func (scq *ScenarioCandidateQuery) Order(o ...scenariocandidate.OrderOption) *ScenarioCandidateQuery {
	scq.order = append(scq.order, o...)
	return scq
}

// QueryScenario chains the current query on the "scenario" edge.
func (scq *ScenarioCandidateQuery) QueryScenario() *ScenarioQuery {
	query := (&ScenarioClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scenariocandidate.Table, scenariocandidate.FieldID, selector),
			sqlgraph.To(scenario.Table, scenario.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, scenariocandidate.ScenarioTable, scenariocandidate.ScenarioColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAttempts chains the current query on the "attempts" edge.
func (scq *ScenarioCandidateQuery) QueryAttempts() *SubmissionAttemptQuery {
	query := (&SubmissionAttemptClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scenariocandidate.Table, scenariocandidate.FieldID, selector),
			sqlgraph.To(submissionattempt.Table, submissionattempt.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, scenariocandidate.AttemptsTable, scenariocandidate.AttemptsColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ScenarioCandidate entity from the query.
// Returns a *NotFoundError when no ScenarioCandidate was found.
func (scq *ScenarioCandidateQuery) First(ctx context.Context) (*ScenarioCandidate, error) {
	nodes, err := scq.Limit(1).All(setContextOp(ctx, scq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{scenariocandidate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (scq *ScenarioCandidateQuery) FirstX(ctx context.Context) *ScenarioCandidate {
	node, err := scq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ScenarioCandidate ID from the query.
// Returns a *NotFoundError when no ScenarioCandidate ID was found.
func (scq *ScenarioCandidateQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = scq.Limit(1).IDs(setContextOp(ctx, scq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{scenariocandidate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (scq *ScenarioCandidateQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := scq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ScenarioCandidate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ScenarioCandidate entity is found.
// Returns a *NotFoundError when no ScenarioCandidate entities are found.
func (scq *ScenarioCandidateQuery) Only(ctx context.Context) (*ScenarioCandidate, error) {
	nodes, err := scq.Limit(2).All(setContextOp(ctx, scq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{scenariocandidate.Label}
	default:
		return nil, &NotSingularError{scenariocandidate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (scq *ScenarioCandidateQuery) OnlyX(ctx context.Context) *ScenarioCandidate {
	node, err := scq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ScenarioCandidate ID in the query.
// Returns a *NotSingularError when more than one ScenarioCandidate ID is found.
// Returns a *NotFoundError when no entities are found.
func (scq *ScenarioCandidateQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = scq.Limit(2).IDs(setContextOp(ctx, scq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{scenariocandidate.Label}
	default:
		err = &NotSingularError{scenariocandidate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (scq *ScenarioCandidateQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := scq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ScenarioCandidates.
func (scq *ScenarioCandidateQuery) All(ctx context.Context) ([]*ScenarioCandidate, error) {
	ctx = setContextOp(ctx, scq.ctx, ent.OpQueryAll)
	if err := scq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ScenarioCandidate, *ScenarioCandidateQuery]()
	return withInterceptors[[]*ScenarioCandidate](ctx, scq, qr, scq.inters)
}

// AllX is like All, but panics if an error occurs.
func (scq *ScenarioCandidateQuery) AllX(ctx context.Context) []*ScenarioCandidate {
	nodes, err := scq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ScenarioCandidate IDs.
func (scq *ScenarioCandidateQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if scq.ctx.Unique == nil && scq.path != nil {
		scq.Unique(true)
	}
	ctx = setContextOp(ctx, scq.ctx, ent.OpQueryIDs)
	if err = scq.Select(scenariocandidate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (scq *ScenarioCandidateQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := scq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (scq *ScenarioCandidateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, scq.ctx, ent.OpQueryCount)
	if err := scq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, scq, querierCount[*ScenarioCandidateQuery](), scq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (scq *ScenarioCandidateQuery) CountX(ctx context.Context) int {
	count, err := scq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (scq *ScenarioCandidateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, scq.ctx, ent.OpQueryExist)
	switch _, err := scq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (scq *ScenarioCandidateQuery) ExistX(ctx context.Context) bool {
	exist, err := scq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ScenarioCandidateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (scq *ScenarioCandidateQuery) Clone() *ScenarioCandidateQuery {
	if scq == nil {
		return nil
	}
	return &ScenarioCandidateQuery{
		config:       scq.config,
		ctx:          scq.ctx.Clone(),
		order:        append([]scenariocandidate.OrderOption{}, scq.order...),
		inters:       append([]Interceptor{}, scq.inters...),
		predicates:   append([]predicate.ScenarioCandidate{}, scq.predicates...),
		withScenario: scq.withScenario.Clone(),
		withAttempts: scq.withAttempts.Clone(),
		// clone intermediate query.
		sql:  scq.sql.Clone(),
		path: scq.path,
	}
}

// WithScenario tells the query-builder to eager-load the nodes that are connected to
// the "scenario" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ScenarioCandidateQuery) WithScenario(opts ...func(*ScenarioQuery)) *ScenarioCandidateQuery {
	query := (&ScenarioClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withScenario = query
	return scq
}

// WithAttempts tells the query-builder to eager-load the nodes that are connected to
// the "attempts" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ScenarioCandidateQuery) WithAttempts(opts ...func(*SubmissionAttemptQuery)) *ScenarioCandidateQuery {
	query := (&SubmissionAttemptClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withAttempts = query
	return scq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ScenarioCandidate.Query().
//		GroupBy(scenariocandidate.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (scq *ScenarioCandidateQuery) GroupBy(field string, fields ...string) *ScenarioCandidateGroupBy {
	scq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ScenarioCandidateGroupBy{build: scq}
	grbuild.flds = &scq.ctx.Fields
	grbuild.label = scenariocandidate.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.ScenarioCandidate.Query().
//		Select(scenariocandidate.FieldCreatedAt).
//		Scan(ctx, &v)
func (scq *ScenarioCandidateQuery) Select(fields ...string) *ScenarioCandidateSelect {
	scq.ctx.Fields = append(scq.ctx.Fields, fields...)
	sbuild := &ScenarioCandidateSelect{ScenarioCandidateQuery: scq}
	sbuild.label = scenariocandidate.Label
	sbuild.flds, sbuild.scan = &scq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ScenarioCandidateSelect configured with the given aggregations.
func (scq *ScenarioCandidateQuery) Aggregate(fns ...AggregateFunc) *ScenarioCandidateSelect {
	return scq.Select().Aggregate(fns...)
}

func (scq *ScenarioCandidateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range scq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, scq); err != nil {
				return err
			}
		}
	}
	for _, f := range scq.ctx.Fields {
		if !scenariocandidate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if scq.path != nil {
		prev, err := scq.path(ctx)
		if err != nil {
			return err
		}
		scq.sql = prev
	}
	return nil
}

func (scq *ScenarioCandidateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ScenarioCandidate, error) {
	var (
		nodes       = []*ScenarioCandidate{}
		_spec       = scq.querySpec()
		loadedTypes = [2]bool{
			scq.withScenario != nil,
			scq.withAttempts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ScenarioCandidate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ScenarioCandidate{config: scq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, scq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := scq.withScenario; query != nil {
		if err := scq.loadScenario(ctx, query, nodes, nil,
			func(n *ScenarioCandidate, e *Scenario) { n.Edges.Scenario = e }); err != nil {
			return nil, err
		}
	}
	if query := scq.withAttempts; query != nil {
		if err := scq.loadAttempts(ctx, query, nodes,
			func(n *ScenarioCandidate) { n.Edges.Attempts = []*SubmissionAttempt{} },
			func(n *ScenarioCandidate, e *SubmissionAttempt) { n.Edges.Attempts = append(n.Edges.Attempts, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (scq *ScenarioCandidateQuery) loadScenario(ctx context.Context, query *ScenarioQuery, nodes []*ScenarioCandidate, init func(*ScenarioCandidate), assign func(*ScenarioCandidate, *Scenario)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*ScenarioCandidate)
	for i := range nodes {
		fk := nodes[i].ScenarioID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(scenario.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "scenario_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (scq *ScenarioCandidateQuery) loadAttempts(ctx context.Context, query *SubmissionAttemptQuery, nodes []*ScenarioCandidate, init func(*ScenarioCandidate), assign func(*ScenarioCandidate, *SubmissionAttempt)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*ScenarioCandidate)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(submissionattempt.FieldScenarioCandidateID)
	}
	query.Where(predicate.SubmissionAttempt(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(scenariocandidate.AttemptsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ScenarioCandidateID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "scenario_candidate_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (scq *ScenarioCandidateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := scq.querySpec()
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	_spec.Node.Columns = scq.ctx.Fields
	if len(scq.ctx.Fields) > 0 {
		_spec.Unique = scq.ctx.Unique != nil && *scq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, scq.driver, _spec)
}

func (scq *ScenarioCandidateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(scenariocandidate.Table, scenariocandidate.Columns, sqlgraph.NewFieldSpec(scenariocandidate.FieldID, field.TypeUint64))
	_spec.From = scq.sql
	if unique := scq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if scq.path != nil {
		_spec.Unique = true
	}
	if fields := scq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scenariocandidate.FieldID)
		for i := range fields {
			if fields[i] != scenariocandidate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if scq.withScenario != nil {
			_spec.Node.AddColumnOnce(scenariocandidate.FieldScenarioID)
		}
	}
	if ps := scq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := scq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := scq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := scq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (scq *ScenarioCandidateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(scq.driver.Dialect())
	t1 := builder.Table(scenariocandidate.Table)
	columns := scq.ctx.Fields
	if len(columns) == 0 {
		columns = scenariocandidate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if scq.sql != nil {
		selector = scq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if scq.ctx.Unique != nil && *scq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range scq.modifiers {
		m(selector)
	}
	for _, p := range scq.predicates {
		p(selector)
	}
	for _, p := range scq.order {
		p(selector)
	}
	if offset := scq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := scq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (scq *ScenarioCandidateQuery) ForUpdate(opts ...sql.LockOption) *ScenarioCandidateQuery {
	if scq.driver.Dialect() == dialect.Postgres {
		scq.Unique(false)
	}
	scq.modifiers = append(scq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return scq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (scq *ScenarioCandidateQuery) ForShare(opts ...sql.LockOption) *ScenarioCandidateQuery {
	if scq.driver.Dialect() == dialect.Postgres {
		scq.Unique(false)
	}
	scq.modifiers = append(scq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return scq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (scq *ScenarioCandidateQuery) Modify(modifiers ...func(s *sql.Selector)) *ScenarioCandidateSelect {
	scq.modifiers = append(scq.modifiers, modifiers...)
	return scq.Select()
}

// ScenarioCandidateGroupBy is the group-by builder for ScenarioCandidate entities.
type ScenarioCandidateGroupBy struct {
	selector
	build *ScenarioCandidateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scgb *ScenarioCandidateGroupBy) Aggregate(fns ...AggregateFunc) *ScenarioCandidateGroupBy {
	scgb.fns = append(scgb.fns, fns...)
	return scgb
}

// Scan applies the selector query and scans the result into the given value.
func (scgb *ScenarioCandidateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scgb.build.ctx, ent.OpQueryGroupBy)
	if err := scgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScenarioCandidateQuery, *ScenarioCandidateGroupBy](ctx, scgb.build, scgb, scgb.build.inters, v)
}

func (scgb *ScenarioCandidateGroupBy) sqlScan(ctx context.Context, root *ScenarioCandidateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(scgb.fns))
	for _, fn := range scgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*scgb.flds)+len(scgb.fns))
		for _, f := range *scgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*scgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ScenarioCandidateSelect is the builder for selecting fields of ScenarioCandidate entities.
type ScenarioCandidateSelect struct {
	*ScenarioCandidateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (scs *ScenarioCandidateSelect) Aggregate(fns ...AggregateFunc) *ScenarioCandidateSelect {
	scs.fns = append(scs.fns, fns...)
	return scs
}

// Scan applies the selector query and scans the result into the given value.
func (scs *ScenarioCandidateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scs.ctx, ent.OpQuerySelect)
	if err := scs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScenarioCandidateQuery, *ScenarioCandidateSelect](ctx, scs.ScenarioCandidateQuery, scs, scs.inters, v)
}

func (scs *ScenarioCandidateSelect) sqlScan(ctx context.Context, root *ScenarioCandidateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(scs.fns))
	for _, fn := range scs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*scs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (scs *ScenarioCandidateSelect) Modify(modifiers ...func(s *sql.Selector)) *ScenarioCandidateSelect {
	scs.modifiers = append(scs.modifiers, modifiers...)
	return scs
}
