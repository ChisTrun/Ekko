// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"ekko/pkg/ent/predicate"
	"ekko/pkg/ent/question"
	"ekko/pkg/ent/scenario"
	"ekko/pkg/ent/scenariocandidate"
	"ekko/pkg/ent/scenariofavorite"
	"ekko/pkg/ent/scenariofield"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScenarioQuery is the builder for querying Scenario entities.
type ScenarioQuery struct {
	config
	ctx            *QueryContext
	order          []scenario.OrderOption
	inters         []Interceptor
	predicates     []predicate.Scenario
	withQuestions  *QuestionQuery
	withCandidates *ScenarioCandidateQuery
	withFavorites  *ScenarioFavoriteQuery
	withField      *ScenarioFieldQuery
	modifiers      []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ScenarioQuery builder.
func (sq *ScenarioQuery) Where(ps ...predicate.Scenario) *ScenarioQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *ScenarioQuery) Limit(limit int) *ScenarioQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *ScenarioQuery) Offset(offset int) *ScenarioQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *ScenarioQuery) Unique(unique bool) *ScenarioQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *ScenarioQuery) Order(o ...scenario.OrderOption) *ScenarioQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryQuestions chains the current query on the "questions" edge.
func (sq *ScenarioQuery) QueryQuestions() *QuestionQuery {
	query := (&QuestionClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scenario.Table, scenario.FieldID, selector),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, scenario.QuestionsTable, scenario.QuestionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidates chains the current query on the "candidates" edge.
func (sq *ScenarioQuery) QueryCandidates() *ScenarioCandidateQuery {
	query := (&ScenarioCandidateClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scenario.Table, scenario.FieldID, selector),
			sqlgraph.To(scenariocandidate.Table, scenariocandidate.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, scenario.CandidatesTable, scenario.CandidatesColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFavorites chains the current query on the "favorites" edge.
func (sq *ScenarioQuery) QueryFavorites() *ScenarioFavoriteQuery {
	query := (&ScenarioFavoriteClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scenario.Table, scenario.FieldID, selector),
			sqlgraph.To(scenariofavorite.Table, scenariofavorite.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, scenario.FavoritesTable, scenario.FavoritesColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryField chains the current query on the "field" edge.
func (sq *ScenarioQuery) QueryField() *ScenarioFieldQuery {
	query := (&ScenarioFieldClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scenario.Table, scenario.FieldID, selector),
			sqlgraph.To(scenariofield.Table, scenariofield.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, scenario.FieldTable, scenario.FieldPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Scenario entity from the query.
// Returns a *NotFoundError when no Scenario was found.
func (sq *ScenarioQuery) First(ctx context.Context) (*Scenario, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{scenario.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ScenarioQuery) FirstX(ctx context.Context) *Scenario {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Scenario ID from the query.
// Returns a *NotFoundError when no Scenario ID was found.
func (sq *ScenarioQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{scenario.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *ScenarioQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Scenario entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Scenario entity is found.
// Returns a *NotFoundError when no Scenario entities are found.
func (sq *ScenarioQuery) Only(ctx context.Context) (*Scenario, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{scenario.Label}
	default:
		return nil, &NotSingularError{scenario.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ScenarioQuery) OnlyX(ctx context.Context) *Scenario {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Scenario ID in the query.
// Returns a *NotSingularError when more than one Scenario ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *ScenarioQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{scenario.Label}
	default:
		err = &NotSingularError{scenario.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *ScenarioQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Scenarios.
func (sq *ScenarioQuery) All(ctx context.Context) ([]*Scenario, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryAll)
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Scenario, *ScenarioQuery]()
	return withInterceptors[[]*Scenario](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *ScenarioQuery) AllX(ctx context.Context) []*Scenario {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Scenario IDs.
func (sq *ScenarioQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryIDs)
	if err = sq.Select(scenario.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ScenarioQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ScenarioQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryCount)
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*ScenarioQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ScenarioQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ScenarioQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryExist)
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ScenarioQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ScenarioQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ScenarioQuery) Clone() *ScenarioQuery {
	if sq == nil {
		return nil
	}
	return &ScenarioQuery{
		config:         sq.config,
		ctx:            sq.ctx.Clone(),
		order:          append([]scenario.OrderOption{}, sq.order...),
		inters:         append([]Interceptor{}, sq.inters...),
		predicates:     append([]predicate.Scenario{}, sq.predicates...),
		withQuestions:  sq.withQuestions.Clone(),
		withCandidates: sq.withCandidates.Clone(),
		withFavorites:  sq.withFavorites.Clone(),
		withField:      sq.withField.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithQuestions tells the query-builder to eager-load the nodes that are connected to
// the "questions" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScenarioQuery) WithQuestions(opts ...func(*QuestionQuery)) *ScenarioQuery {
	query := (&QuestionClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withQuestions = query
	return sq
}

// WithCandidates tells the query-builder to eager-load the nodes that are connected to
// the "candidates" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScenarioQuery) WithCandidates(opts ...func(*ScenarioCandidateQuery)) *ScenarioQuery {
	query := (&ScenarioCandidateClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withCandidates = query
	return sq
}

// WithFavorites tells the query-builder to eager-load the nodes that are connected to
// the "favorites" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScenarioQuery) WithFavorites(opts ...func(*ScenarioFavoriteQuery)) *ScenarioQuery {
	query := (&ScenarioFavoriteClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withFavorites = query
	return sq
}

// WithField tells the query-builder to eager-load the nodes that are connected to
// the "field" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScenarioQuery) WithField(opts ...func(*ScenarioFieldQuery)) *ScenarioQuery {
	query := (&ScenarioFieldClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withField = query
	return sq
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
//	client.Scenario.Query().
//		GroupBy(scenario.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *ScenarioQuery) GroupBy(field string, fields ...string) *ScenarioGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ScenarioGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = scenario.Label
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
//	client.Scenario.Query().
//		Select(scenario.FieldCreatedAt).
//		Scan(ctx, &v)
func (sq *ScenarioQuery) Select(fields ...string) *ScenarioSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &ScenarioSelect{ScenarioQuery: sq}
	sbuild.label = scenario.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ScenarioSelect configured with the given aggregations.
func (sq *ScenarioQuery) Aggregate(fns ...AggregateFunc) *ScenarioSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *ScenarioQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !scenario.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *ScenarioQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Scenario, error) {
	var (
		nodes       = []*Scenario{}
		_spec       = sq.querySpec()
		loadedTypes = [4]bool{
			sq.withQuestions != nil,
			sq.withCandidates != nil,
			sq.withFavorites != nil,
			sq.withField != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Scenario).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Scenario{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withQuestions; query != nil {
		if err := sq.loadQuestions(ctx, query, nodes,
			func(n *Scenario) { n.Edges.Questions = []*Question{} },
			func(n *Scenario, e *Question) { n.Edges.Questions = append(n.Edges.Questions, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withCandidates; query != nil {
		if err := sq.loadCandidates(ctx, query, nodes,
			func(n *Scenario) { n.Edges.Candidates = []*ScenarioCandidate{} },
			func(n *Scenario, e *ScenarioCandidate) { n.Edges.Candidates = append(n.Edges.Candidates, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withFavorites; query != nil {
		if err := sq.loadFavorites(ctx, query, nodes,
			func(n *Scenario) { n.Edges.Favorites = []*ScenarioFavorite{} },
			func(n *Scenario, e *ScenarioFavorite) { n.Edges.Favorites = append(n.Edges.Favorites, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withField; query != nil {
		if err := sq.loadField(ctx, query, nodes,
			func(n *Scenario) { n.Edges.Field = []*ScenarioField{} },
			func(n *Scenario, e *ScenarioField) { n.Edges.Field = append(n.Edges.Field, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *ScenarioQuery) loadQuestions(ctx context.Context, query *QuestionQuery, nodes []*Scenario, init func(*Scenario), assign func(*Scenario, *Question)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Scenario)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(question.FieldScenarioID)
	}
	query.Where(predicate.Question(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(scenario.QuestionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ScenarioID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "scenario_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *ScenarioQuery) loadCandidates(ctx context.Context, query *ScenarioCandidateQuery, nodes []*Scenario, init func(*Scenario), assign func(*Scenario, *ScenarioCandidate)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Scenario)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(scenariocandidate.FieldScenarioID)
	}
	query.Where(predicate.ScenarioCandidate(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(scenario.CandidatesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ScenarioID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "scenario_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *ScenarioQuery) loadFavorites(ctx context.Context, query *ScenarioFavoriteQuery, nodes []*Scenario, init func(*Scenario), assign func(*Scenario, *ScenarioFavorite)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Scenario)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(scenariofavorite.FieldScenarioID)
	}
	query.Where(predicate.ScenarioFavorite(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(scenario.FavoritesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ScenarioID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "scenario_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *ScenarioQuery) loadField(ctx context.Context, query *ScenarioFieldQuery, nodes []*Scenario, init func(*Scenario), assign func(*Scenario, *ScenarioField)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uint64]*Scenario)
	nids := make(map[uint64]map[*Scenario]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(scenario.FieldTable)
		s.Join(joinT).On(s.C(scenariofield.FieldID), joinT.C(scenario.FieldPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(scenario.FieldPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(scenario.FieldPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := uint64(values[0].(*sql.NullInt64).Int64)
				inValue := uint64(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Scenario]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*ScenarioField](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "field" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (sq *ScenarioQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ScenarioQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(scenario.Table, scenario.Columns, sqlgraph.NewFieldSpec(scenario.FieldID, field.TypeUint64))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scenario.FieldID)
		for i := range fields {
			if fields[i] != scenario.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ScenarioQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(scenario.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = scenario.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range sq.modifiers {
		m(selector)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (sq *ScenarioQuery) ForUpdate(opts ...sql.LockOption) *ScenarioQuery {
	if sq.driver.Dialect() == dialect.Postgres {
		sq.Unique(false)
	}
	sq.modifiers = append(sq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return sq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (sq *ScenarioQuery) ForShare(opts ...sql.LockOption) *ScenarioQuery {
	if sq.driver.Dialect() == dialect.Postgres {
		sq.Unique(false)
	}
	sq.modifiers = append(sq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return sq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sq *ScenarioQuery) Modify(modifiers ...func(s *sql.Selector)) *ScenarioSelect {
	sq.modifiers = append(sq.modifiers, modifiers...)
	return sq.Select()
}

// ScenarioGroupBy is the group-by builder for Scenario entities.
type ScenarioGroupBy struct {
	selector
	build *ScenarioQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ScenarioGroupBy) Aggregate(fns ...AggregateFunc) *ScenarioGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *ScenarioGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, ent.OpQueryGroupBy)
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScenarioQuery, *ScenarioGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *ScenarioGroupBy) sqlScan(ctx context.Context, root *ScenarioQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ScenarioSelect is the builder for selecting fields of Scenario entities.
type ScenarioSelect struct {
	*ScenarioQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *ScenarioSelect) Aggregate(fns ...AggregateFunc) *ScenarioSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *ScenarioSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, ent.OpQuerySelect)
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScenarioQuery, *ScenarioSelect](ctx, ss.ScenarioQuery, ss, ss.inters, v)
}

func (ss *ScenarioSelect) sqlScan(ctx context.Context, root *ScenarioQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ss *ScenarioSelect) Modify(modifiers ...func(s *sql.Selector)) *ScenarioSelect {
	ss.modifiers = append(ss.modifiers, modifiers...)
	return ss
}
