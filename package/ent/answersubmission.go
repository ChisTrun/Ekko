// Code generated by ent, DO NOT EDIT.

package ent

import (
	ekko "ekko/api"
	"ekko/package/ent/answersubmission"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// AnswerSubmission is the model entity for the AnswerSubmission schema.
type AnswerSubmission struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// SubmissionAttemptID holds the value of the "submission_attempt_id" field.
	SubmissionAttemptID uint64 `json:"submission_attempt_id,omitempty"`
	// QuestionID holds the value of the "question_id" field.
	QuestionID uint64 `json:"question_id,omitempty"`
	// Answer holds the value of the "answer" field.
	Answer string `json:"answer,omitempty"`
	// Relevance holds the value of the "relevance" field.
	Relevance float64 `json:"relevance,omitempty"`
	// ClarityCompleteness holds the value of the "clarity_completeness" field.
	ClarityCompleteness float64 `json:"clarity_completeness,omitempty"`
	// Accuracy holds the value of the "accuracy" field.
	Accuracy float64 `json:"accuracy,omitempty"`
	// Overall holds the value of the "overall" field.
	Overall float64 `json:"overall,omitempty"`
	// Status holds the value of the "status" field.
	Status       ekko.SubmissionStatus `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AnswerSubmission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case answersubmission.FieldRelevance, answersubmission.FieldClarityCompleteness, answersubmission.FieldAccuracy, answersubmission.FieldOverall:
			values[i] = new(sql.NullFloat64)
		case answersubmission.FieldID, answersubmission.FieldSubmissionAttemptID, answersubmission.FieldQuestionID, answersubmission.FieldStatus:
			values[i] = new(sql.NullInt64)
		case answersubmission.FieldAnswer:
			values[i] = new(sql.NullString)
		case answersubmission.FieldCreatedAt, answersubmission.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AnswerSubmission fields.
func (as *AnswerSubmission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case answersubmission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			as.ID = uint64(value.Int64)
		case answersubmission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				as.CreatedAt = value.Time
			}
		case answersubmission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				as.UpdatedAt = value.Time
			}
		case answersubmission.FieldSubmissionAttemptID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field submission_attempt_id", values[i])
			} else if value.Valid {
				as.SubmissionAttemptID = uint64(value.Int64)
			}
		case answersubmission.FieldQuestionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field question_id", values[i])
			} else if value.Valid {
				as.QuestionID = uint64(value.Int64)
			}
		case answersubmission.FieldAnswer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field answer", values[i])
			} else if value.Valid {
				as.Answer = value.String
			}
		case answersubmission.FieldRelevance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field relevance", values[i])
			} else if value.Valid {
				as.Relevance = value.Float64
			}
		case answersubmission.FieldClarityCompleteness:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field clarity_completeness", values[i])
			} else if value.Valid {
				as.ClarityCompleteness = value.Float64
			}
		case answersubmission.FieldAccuracy:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field accuracy", values[i])
			} else if value.Valid {
				as.Accuracy = value.Float64
			}
		case answersubmission.FieldOverall:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field overall", values[i])
			} else if value.Valid {
				as.Overall = value.Float64
			}
		case answersubmission.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				as.Status = ekko.SubmissionStatus(value.Int64)
			}
		default:
			as.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AnswerSubmission.
// This includes values selected through modifiers, order, etc.
func (as *AnswerSubmission) Value(name string) (ent.Value, error) {
	return as.selectValues.Get(name)
}

// Update returns a builder for updating this AnswerSubmission.
// Note that you need to call AnswerSubmission.Unwrap() before calling this method if this AnswerSubmission
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *AnswerSubmission) Update() *AnswerSubmissionUpdateOne {
	return NewAnswerSubmissionClient(as.config).UpdateOne(as)
}

// Unwrap unwraps the AnswerSubmission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *AnswerSubmission) Unwrap() *AnswerSubmission {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("ent: AnswerSubmission is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *AnswerSubmission) String() string {
	var builder strings.Builder
	builder.WriteString("AnswerSubmission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	builder.WriteString("created_at=")
	builder.WriteString(as.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(as.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("submission_attempt_id=")
	builder.WriteString(fmt.Sprintf("%v", as.SubmissionAttemptID))
	builder.WriteString(", ")
	builder.WriteString("question_id=")
	builder.WriteString(fmt.Sprintf("%v", as.QuestionID))
	builder.WriteString(", ")
	builder.WriteString("answer=")
	builder.WriteString(as.Answer)
	builder.WriteString(", ")
	builder.WriteString("relevance=")
	builder.WriteString(fmt.Sprintf("%v", as.Relevance))
	builder.WriteString(", ")
	builder.WriteString("clarity_completeness=")
	builder.WriteString(fmt.Sprintf("%v", as.ClarityCompleteness))
	builder.WriteString(", ")
	builder.WriteString("accuracy=")
	builder.WriteString(fmt.Sprintf("%v", as.Accuracy))
	builder.WriteString(", ")
	builder.WriteString("overall=")
	builder.WriteString(fmt.Sprintf("%v", as.Overall))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", as.Status))
	builder.WriteByte(')')
	return builder.String()
}

// AnswerSubmissions is a parsable slice of AnswerSubmission.
type AnswerSubmissions []*AnswerSubmission
