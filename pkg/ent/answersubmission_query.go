// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ekko/pkg/ent/answersubmission"
	"ekko/pkg/ent/predicate"
	"ekko/pkg/ent/submissionattempt"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AnswerSubmissionQuery is the builder for querying AnswerSubmission entities.
type AnswerSubmissionQuery struct {
	config
	ctx                   *QueryContext
	order                 []answersubmission.OrderOption
	inters                []Interceptor
	predicates            []predicate.AnswerSubmission
	withSubmissionAttempt *SubmissionAttemptQuery
	modifiers             []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AnswerSubmissionQuery builder.
func (asq *AnswerSubmissionQuery) Where(ps ...predicate.AnswerSubmission) *AnswerSubmissionQuery {
	asq.predicates = append(asq.predicates, ps...)
	return asq
}

// Limit the number of records to be returned by this query.
func (asq *AnswerSubmissionQuery) Limit(limit int) *AnswerSubmissionQuery {
	asq.ctx.Limit = &limit
	return asq
}

// Offset to start from.
func (asq *AnswerSubmissionQuery) Offset(offset int) *AnswerSubmissionQuery {
	asq.ctx.Offset = &offset
	return asq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (asq *AnswerSubmissionQuery) Unique(unique bool) *AnswerSubmissionQuery {
	asq.ctx.Unique = &unique
	return asq
}

// Order specifies how the records should be ordered.
func (asq *AnswerSubmissionQuery) Order(o ...answersubmission.OrderOption) *AnswerSubmissionQuery {
	asq.order = append(asq.order, o...)
	return asq
}

// QuerySubmissionAttempt chains the current query on the "submission_attempt" edge.
func (asq *AnswerSubmissionQuery) QuerySubmissionAttempt() *SubmissionAttemptQuery {
	query := (&SubmissionAttemptClient{config: asq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := asq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := asq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(answersubmission.Table, answersubmission.FieldID, selector),
			sqlgraph.To(submissionattempt.Table, submissionattempt.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, answersubmission.SubmissionAttemptTable, answersubmission.SubmissionAttemptColumn),
		)
		fromU = sqlgraph.SetNeighbors(asq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AnswerSubmission entity from the query.
// Returns a *NotFoundError when no AnswerSubmission was found.
func (asq *AnswerSubmissionQuery) First(ctx context.Context) (*AnswerSubmission, error) {
	nodes, err := asq.Limit(1).All(setContextOp(ctx, asq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{answersubmission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (asq *AnswerSubmissionQuery) FirstX(ctx context.Context) *AnswerSubmission {
	node, err := asq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AnswerSubmission ID from the query.
// Returns a *NotFoundError when no AnswerSubmission ID was found.
func (asq *AnswerSubmissionQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = asq.Limit(1).IDs(setContextOp(ctx, asq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{answersubmission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (asq *AnswerSubmissionQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := asq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AnswerSubmission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AnswerSubmission entity is found.
// Returns a *NotFoundError when no AnswerSubmission entities are found.
func (asq *AnswerSubmissionQuery) Only(ctx context.Context) (*AnswerSubmission, error) {
	nodes, err := asq.Limit(2).All(setContextOp(ctx, asq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{answersubmission.Label}
	default:
		return nil, &NotSingularError{answersubmission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (asq *AnswerSubmissionQuery) OnlyX(ctx context.Context) *AnswerSubmission {
	node, err := asq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AnswerSubmission ID in the query.
// Returns a *NotSingularError when more than one AnswerSubmission ID is found.
// Returns a *NotFoundError when no entities are found.
func (asq *AnswerSubmissionQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = asq.Limit(2).IDs(setContextOp(ctx, asq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{answersubmission.Label}
	default:
		err = &NotSingularError{answersubmission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (asq *AnswerSubmissionQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := asq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AnswerSubmissions.
func (asq *AnswerSubmissionQuery) All(ctx context.Context) ([]*AnswerSubmission, error) {
	ctx = setContextOp(ctx, asq.ctx, ent.OpQueryAll)
	if err := asq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AnswerSubmission, *AnswerSubmissionQuery]()
	return withInterceptors[[]*AnswerSubmission](ctx, asq, qr, asq.inters)
}

// AllX is like All, but panics if an error occurs.
func (asq *AnswerSubmissionQuery) AllX(ctx context.Context) []*AnswerSubmission {
	nodes, err := asq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AnswerSubmission IDs.
func (asq *AnswerSubmissionQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if asq.ctx.Unique == nil && asq.path != nil {
		asq.Unique(true)
	}
	ctx = setContextOp(ctx, asq.ctx, ent.OpQueryIDs)
	if err = asq.Select(answersubmission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (asq *AnswerSubmissionQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := asq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (asq *AnswerSubmissionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, asq.ctx, ent.OpQueryCount)
	if err := asq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, asq, querierCount[*AnswerSubmissionQuery](), asq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (asq *AnswerSubmissionQuery) CountX(ctx context.Context) int {
	count, err := asq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (asq *AnswerSubmissionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, asq.ctx, ent.OpQueryExist)
	switch _, err := asq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (asq *AnswerSubmissionQuery) ExistX(ctx context.Context) bool {
	exist, err := asq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AnswerSubmissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (asq *AnswerSubmissionQuery) Clone() *AnswerSubmissionQuery {
	if asq == nil {
		return nil
	}
	return &AnswerSubmissionQuery{
		config:                asq.config,
		ctx:                   asq.ctx.Clone(),
		order:                 append([]answersubmission.OrderOption{}, asq.order...),
		inters:                append([]Interceptor{}, asq.inters...),
		predicates:            append([]predicate.AnswerSubmission{}, asq.predicates...),
		withSubmissionAttempt: asq.withSubmissionAttempt.Clone(),
		// clone intermediate query.
		sql:  asq.sql.Clone(),
		path: asq.path,
	}
}

// WithSubmissionAttempt tells the query-builder to eager-load the nodes that are connected to
// the "submission_attempt" edge. The optional arguments are used to configure the query builder of the edge.
func (asq *AnswerSubmissionQuery) WithSubmissionAttempt(opts ...func(*SubmissionAttemptQuery)) *AnswerSubmissionQuery {
	query := (&SubmissionAttemptClient{config: asq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	asq.withSubmissionAttempt = query
	return asq
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
//	client.AnswerSubmission.Query().
//		GroupBy(answersubmission.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (asq *AnswerSubmissionQuery) GroupBy(field string, fields ...string) *AnswerSubmissionGroupBy {
	asq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AnswerSubmissionGroupBy{build: asq}
	grbuild.flds = &asq.ctx.Fields
	grbuild.label = answersubmission.Label
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
//	client.AnswerSubmission.Query().
//		Select(answersubmission.FieldCreatedAt).
//		Scan(ctx, &v)
func (asq *AnswerSubmissionQuery) Select(fields ...string) *AnswerSubmissionSelect {
	asq.ctx.Fields = append(asq.ctx.Fields, fields...)
	sbuild := &AnswerSubmissionSelect{AnswerSubmissionQuery: asq}
	sbuild.label = answersubmission.Label
	sbuild.flds, sbuild.scan = &asq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AnswerSubmissionSelect configured with the given aggregations.
func (asq *AnswerSubmissionQuery) Aggregate(fns ...AggregateFunc) *AnswerSubmissionSelect {
	return asq.Select().Aggregate(fns...)
}

func (asq *AnswerSubmissionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range asq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, asq); err != nil {
				return err
			}
		}
	}
	for _, f := range asq.ctx.Fields {
		if !answersubmission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if asq.path != nil {
		prev, err := asq.path(ctx)
		if err != nil {
			return err
		}
		asq.sql = prev
	}
	return nil
}

func (asq *AnswerSubmissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AnswerSubmission, error) {
	var (
		nodes       = []*AnswerSubmission{}
		_spec       = asq.querySpec()
		loadedTypes = [1]bool{
			asq.withSubmissionAttempt != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AnswerSubmission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AnswerSubmission{config: asq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(asq.modifiers) > 0 {
		_spec.Modifiers = asq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, asq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := asq.withSubmissionAttempt; query != nil {
		if err := asq.loadSubmissionAttempt(ctx, query, nodes, nil,
			func(n *AnswerSubmission, e *SubmissionAttempt) { n.Edges.SubmissionAttempt = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (asq *AnswerSubmissionQuery) loadSubmissionAttempt(ctx context.Context, query *SubmissionAttemptQuery, nodes []*AnswerSubmission, init func(*AnswerSubmission), assign func(*AnswerSubmission, *SubmissionAttempt)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*AnswerSubmission)
	for i := range nodes {
		fk := nodes[i].SubmissionAttemptID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(submissionattempt.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "submission_attempt_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (asq *AnswerSubmissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := asq.querySpec()
	if len(asq.modifiers) > 0 {
		_spec.Modifiers = asq.modifiers
	}
	_spec.Node.Columns = asq.ctx.Fields
	if len(asq.ctx.Fields) > 0 {
		_spec.Unique = asq.ctx.Unique != nil && *asq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, asq.driver, _spec)
}

func (asq *AnswerSubmissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(answersubmission.Table, answersubmission.Columns, sqlgraph.NewFieldSpec(answersubmission.FieldID, field.TypeUint64))
	_spec.From = asq.sql
	if unique := asq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if asq.path != nil {
		_spec.Unique = true
	}
	if fields := asq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, answersubmission.FieldID)
		for i := range fields {
			if fields[i] != answersubmission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if asq.withSubmissionAttempt != nil {
			_spec.Node.AddColumnOnce(answersubmission.FieldSubmissionAttemptID)
		}
	}
	if ps := asq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := asq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := asq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := asq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (asq *AnswerSubmissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(asq.driver.Dialect())
	t1 := builder.Table(answersubmission.Table)
	columns := asq.ctx.Fields
	if len(columns) == 0 {
		columns = answersubmission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if asq.sql != nil {
		selector = asq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if asq.ctx.Unique != nil && *asq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range asq.modifiers {
		m(selector)
	}
	for _, p := range asq.predicates {
		p(selector)
	}
	for _, p := range asq.order {
		p(selector)
	}
	if offset := asq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := asq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (asq *AnswerSubmissionQuery) ForUpdate(opts ...sql.LockOption) *AnswerSubmissionQuery {
	if asq.driver.Dialect() == dialect.Postgres {
		asq.Unique(false)
	}
	asq.modifiers = append(asq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return asq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (asq *AnswerSubmissionQuery) ForShare(opts ...sql.LockOption) *AnswerSubmissionQuery {
	if asq.driver.Dialect() == dialect.Postgres {
		asq.Unique(false)
	}
	asq.modifiers = append(asq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return asq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (asq *AnswerSubmissionQuery) Modify(modifiers ...func(s *sql.Selector)) *AnswerSubmissionSelect {
	asq.modifiers = append(asq.modifiers, modifiers...)
	return asq.Select()
}

// AnswerSubmissionGroupBy is the group-by builder for AnswerSubmission entities.
type AnswerSubmissionGroupBy struct {
	selector
	build *AnswerSubmissionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (asgb *AnswerSubmissionGroupBy) Aggregate(fns ...AggregateFunc) *AnswerSubmissionGroupBy {
	asgb.fns = append(asgb.fns, fns...)
	return asgb
}

// Scan applies the selector query and scans the result into the given value.
func (asgb *AnswerSubmissionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, asgb.build.ctx, ent.OpQueryGroupBy)
	if err := asgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnswerSubmissionQuery, *AnswerSubmissionGroupBy](ctx, asgb.build, asgb, asgb.build.inters, v)
}

func (asgb *AnswerSubmissionGroupBy) sqlScan(ctx context.Context, root *AnswerSubmissionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(asgb.fns))
	for _, fn := range asgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*asgb.flds)+len(asgb.fns))
		for _, f := range *asgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*asgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := asgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AnswerSubmissionSelect is the builder for selecting fields of AnswerSubmission entities.
type AnswerSubmissionSelect struct {
	*AnswerSubmissionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ass *AnswerSubmissionSelect) Aggregate(fns ...AggregateFunc) *AnswerSubmissionSelect {
	ass.fns = append(ass.fns, fns...)
	return ass
}

// Scan applies the selector query and scans the result into the given value.
func (ass *AnswerSubmissionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ass.ctx, ent.OpQuerySelect)
	if err := ass.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnswerSubmissionQuery, *AnswerSubmissionSelect](ctx, ass.AnswerSubmissionQuery, ass, ass.inters, v)
}

func (ass *AnswerSubmissionSelect) sqlScan(ctx context.Context, root *AnswerSubmissionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ass.fns))
	for _, fn := range ass.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ass.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ass.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ass *AnswerSubmissionSelect) Modify(modifiers ...func(s *sql.Selector)) *AnswerSubmissionSelect {
	ass.modifiers = append(ass.modifiers, modifiers...)
	return ass
}
