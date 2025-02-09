// Code generated by ent, DO NOT EDIT.

package question

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the question type in the database.
	Label = "question"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSentenceID holds the string denoting the sentence_id field in the database.
	FieldSentenceID = "sentence_id"
	// FieldCriteria holds the string denoting the criteria field in the database.
	FieldCriteria = "criteria"
	// FieldHint holds the string denoting the hint field in the database.
	FieldHint = "hint"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// EdgeScenario holds the string denoting the scenario edge name in mutations.
	EdgeScenario = "scenario"
	// Table holds the table name of the question in the database.
	Table = "questions"
	// ScenarioTable is the table that holds the scenario relation/edge.
	ScenarioTable = "questions"
	// ScenarioInverseTable is the table name for the Scenario entity.
	// It exists in this package in order to avoid circular dependency with the "scenario" package.
	ScenarioInverseTable = "scenarios"
	// ScenarioColumn is the table column denoting the scenario relation/edge.
	ScenarioColumn = "sentence_id"
)

// Columns holds all SQL columns for question fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSentenceID,
	FieldCriteria,
	FieldHint,
	FieldContent,
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

// OrderOption defines the ordering options for the Question queries.
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

// BySentenceID orders the results by the sentence_id field.
func BySentenceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSentenceID, opts...).ToFunc()
}

// ByCriteria orders the results by the criteria field.
func ByCriteria(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCriteria, opts...).ToFunc()
}

// ByHint orders the results by the hint field.
func ByHint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHint, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByScenarioField orders the results by scenario field.
func ByScenarioField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newScenarioStep(), sql.OrderByField(field, opts...))
	}
}
func newScenarioStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScenarioInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ScenarioTable, ScenarioColumn),
	)
}
