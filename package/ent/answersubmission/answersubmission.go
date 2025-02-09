// Code generated by ent, DO NOT EDIT.

package answersubmission

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the answersubmission type in the database.
	Label = "answer_submission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSubmissionAttemptID holds the string denoting the submission_attempt_id field in the database.
	FieldSubmissionAttemptID = "submission_attempt_id"
	// FieldQuestionID holds the string denoting the question_id field in the database.
	FieldQuestionID = "question_id"
	// FieldAnswer holds the string denoting the answer field in the database.
	FieldAnswer = "answer"
	// FieldRelevance holds the string denoting the relevance field in the database.
	FieldRelevance = "relevance"
	// FieldClarityCompleteness holds the string denoting the clarity_completeness field in the database.
	FieldClarityCompleteness = "clarity_completeness"
	// FieldAccuracy holds the string denoting the accuracy field in the database.
	FieldAccuracy = "accuracy"
	// FieldOverall holds the string denoting the overall field in the database.
	FieldOverall = "overall"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the answersubmission in the database.
	Table = "answer_submissions"
)

// Columns holds all SQL columns for answersubmission fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSubmissionAttemptID,
	FieldQuestionID,
	FieldAnswer,
	FieldRelevance,
	FieldClarityCompleteness,
	FieldAccuracy,
	FieldOverall,
	FieldStatus,
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

// OrderOption defines the ordering options for the AnswerSubmission queries.
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

// BySubmissionAttemptID orders the results by the submission_attempt_id field.
func BySubmissionAttemptID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubmissionAttemptID, opts...).ToFunc()
}

// ByQuestionID orders the results by the question_id field.
func ByQuestionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuestionID, opts...).ToFunc()
}

// ByAnswer orders the results by the answer field.
func ByAnswer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAnswer, opts...).ToFunc()
}

// ByRelevance orders the results by the relevance field.
func ByRelevance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelevance, opts...).ToFunc()
}

// ByClarityCompleteness orders the results by the clarity_completeness field.
func ByClarityCompleteness(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClarityCompleteness, opts...).ToFunc()
}

// ByAccuracy orders the results by the accuracy field.
func ByAccuracy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccuracy, opts...).ToFunc()
}

// ByOverall orders the results by the overall field.
func ByOverall(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOverall, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}
