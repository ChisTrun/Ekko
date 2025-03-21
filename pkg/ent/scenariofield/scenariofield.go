// Code generated by ent, DO NOT EDIT.

package scenariofield

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the scenariofield type in the database.
	Label = "scenario_field"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeSenarios holds the string denoting the senarios edge name in mutations.
	EdgeSenarios = "senarios"
	// Table holds the table name of the scenariofield in the database.
	Table = "scenario_fields"
	// SenariosTable is the table that holds the senarios relation/edge. The primary key declared below.
	SenariosTable = "scenario_field_senarios"
	// SenariosInverseTable is the table name for the Scenario entity.
	// It exists in this package in order to avoid circular dependency with the "scenario" package.
	SenariosInverseTable = "scenarios"
)

// Columns holds all SQL columns for scenariofield fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
}

var (
	// SenariosPrimaryKey and SenariosColumn2 are the table columns denoting the
	// primary key for the senarios relation (M2M).
	SenariosPrimaryKey = []string{"scenario_field_id", "scenario_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the ScenarioField queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySenariosCount orders the results by senarios count.
func BySenariosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSenariosStep(), opts...)
	}
}

// BySenarios orders the results by senarios terms.
func BySenarios(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSenariosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSenariosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SenariosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SenariosTable, SenariosPrimaryKey...),
	)
}
