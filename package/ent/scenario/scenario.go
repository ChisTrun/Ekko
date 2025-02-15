// Code generated by ent, DO NOT EDIT.

package scenario

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the scenario type in the database.
	Label = "scenario"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldBmID holds the string denoting the bm_id field in the database.
	FieldBmID = "bm_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldRating holds the string denoting the rating field in the database.
	FieldRating = "rating"
	// FieldParticipants holds the string denoting the participants field in the database.
	FieldParticipants = "participants"
	// EdgeQuestions holds the string denoting the questions edge name in mutations.
	EdgeQuestions = "questions"
	// EdgeCandidates holds the string denoting the candidates edge name in mutations.
	EdgeCandidates = "candidates"
	// EdgeFavorites holds the string denoting the favorites edge name in mutations.
	EdgeFavorites = "favorites"
	// EdgeField holds the string denoting the field edge name in mutations.
	EdgeField = "field"
	// Table holds the table name of the scenario in the database.
	Table = "scenarios"
	// QuestionsTable is the table that holds the questions relation/edge.
	QuestionsTable = "questions"
	// QuestionsInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionsInverseTable = "questions"
	// QuestionsColumn is the table column denoting the questions relation/edge.
	QuestionsColumn = "scenario_id"
	// CandidatesTable is the table that holds the candidates relation/edge.
	CandidatesTable = "scenario_candidates"
	// CandidatesInverseTable is the table name for the ScenarioCandidate entity.
	// It exists in this package in order to avoid circular dependency with the "scenariocandidate" package.
	CandidatesInverseTable = "scenario_candidates"
	// CandidatesColumn is the table column denoting the candidates relation/edge.
	CandidatesColumn = "scenario_id"
	// FavoritesTable is the table that holds the favorites relation/edge.
	FavoritesTable = "scenario_favorites"
	// FavoritesInverseTable is the table name for the ScenarioFavorite entity.
	// It exists in this package in order to avoid circular dependency with the "scenariofavorite" package.
	FavoritesInverseTable = "scenario_favorites"
	// FavoritesColumn is the table column denoting the favorites relation/edge.
	FavoritesColumn = "scenario_id"
	// FieldTable is the table that holds the field relation/edge.
	FieldTable = "scenario_fields"
	// FieldInverseTable is the table name for the ScenarioField entity.
	// It exists in this package in order to avoid circular dependency with the "scenariofield" package.
	FieldInverseTable = "scenario_fields"
	// FieldColumn is the table column denoting the field relation/edge.
	FieldColumn = "scenario_field_senarios"
)

// Columns holds all SQL columns for scenario fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldBmID,
	FieldName,
	FieldDescription,
	FieldRating,
	FieldParticipants,
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
	// DefaultRating holds the default value on creation for the "rating" field.
	DefaultRating float64
	// DefaultParticipants holds the default value on creation for the "participants" field.
	DefaultParticipants int32
)

// OrderOption defines the ordering options for the Scenario queries.
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

// ByBmID orders the results by the bm_id field.
func ByBmID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBmID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByRating orders the results by the rating field.
func ByRating(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRating, opts...).ToFunc()
}

// ByParticipants orders the results by the participants field.
func ByParticipants(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParticipants, opts...).ToFunc()
}

// ByQuestionsCount orders the results by questions count.
func ByQuestionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newQuestionsStep(), opts...)
	}
}

// ByQuestions orders the results by questions terms.
func ByQuestions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCandidatesCount orders the results by candidates count.
func ByCandidatesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCandidatesStep(), opts...)
	}
}

// ByCandidates orders the results by candidates terms.
func ByCandidates(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCandidatesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFavoritesCount orders the results by favorites count.
func ByFavoritesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFavoritesStep(), opts...)
	}
}

// ByFavorites orders the results by favorites terms.
func ByFavorites(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFavoritesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFieldCount orders the results by field count.
func ByFieldCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFieldStep(), opts...)
	}
}

// ByField orders the results by field terms.
func ByField(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFieldStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newQuestionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, QuestionsTable, QuestionsColumn),
	)
}
func newCandidatesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CandidatesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CandidatesTable, CandidatesColumn),
	)
}
func newFavoritesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FavoritesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FavoritesTable, FavoritesColumn),
	)
}
func newFieldStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FieldInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, FieldTable, FieldColumn),
	)
}
