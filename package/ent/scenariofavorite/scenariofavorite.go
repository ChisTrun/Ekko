// Code generated by ent, DO NOT EDIT.

package scenariofavorite

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the scenariofavorite type in the database.
	Label = "scenario_favorite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldScenarioID holds the string denoting the scenario_id field in the database.
	FieldScenarioID = "scenario_id"
	// EdgeSenario holds the string denoting the senario edge name in mutations.
	EdgeSenario = "senario"
	// Table holds the table name of the scenariofavorite in the database.
	Table = "scenario_favorites"
	// SenarioTable is the table that holds the senario relation/edge.
	SenarioTable = "scenario_favorites"
	// SenarioInverseTable is the table name for the Scenario entity.
	// It exists in this package in order to avoid circular dependency with the "scenario" package.
	SenarioInverseTable = "scenarios"
	// SenarioColumn is the table column denoting the senario relation/edge.
	SenarioColumn = "scenario_id"
)

// Columns holds all SQL columns for scenariofavorite fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUserID,
	FieldScenarioID,
}

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

// OrderOption defines the ordering options for the ScenarioFavorite queries.
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

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByScenarioID orders the results by the scenario_id field.
func ByScenarioID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScenarioID, opts...).ToFunc()
}

// BySenarioField orders the results by senario field.
func BySenarioField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSenarioStep(), sql.OrderByField(field, opts...))
	}
}
func newSenarioStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SenarioInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SenarioTable, SenarioColumn),
	)
}
