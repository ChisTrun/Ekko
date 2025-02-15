// Code generated by ent, DO NOT EDIT.

package submissionattempt

import (
	"ekko/package/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldEQ(FieldUpdatedAt, v))
}

// ScenarioCandidateID applies equality check predicate on the "scenario_candidate_id" field. It's identical to ScenarioCandidateIDEQ.
func ScenarioCandidateID(v uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldEQ(FieldScenarioCandidateID, v))
}

// AttemptNumber applies equality check predicate on the "attempt_number" field. It's identical to AttemptNumberEQ.
func AttemptNumber(v int32) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldEQ(FieldAttemptNumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldLTE(FieldUpdatedAt, v))
}

// ScenarioCandidateIDEQ applies the EQ predicate on the "scenario_candidate_id" field.
func ScenarioCandidateIDEQ(v uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldEQ(FieldScenarioCandidateID, v))
}

// ScenarioCandidateIDNEQ applies the NEQ predicate on the "scenario_candidate_id" field.
func ScenarioCandidateIDNEQ(v uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldNEQ(FieldScenarioCandidateID, v))
}

// ScenarioCandidateIDIn applies the In predicate on the "scenario_candidate_id" field.
func ScenarioCandidateIDIn(vs ...uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldIn(FieldScenarioCandidateID, vs...))
}

// ScenarioCandidateIDNotIn applies the NotIn predicate on the "scenario_candidate_id" field.
func ScenarioCandidateIDNotIn(vs ...uint64) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldNotIn(FieldScenarioCandidateID, vs...))
}

// AttemptNumberEQ applies the EQ predicate on the "attempt_number" field.
func AttemptNumberEQ(v int32) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldEQ(FieldAttemptNumber, v))
}

// AttemptNumberNEQ applies the NEQ predicate on the "attempt_number" field.
func AttemptNumberNEQ(v int32) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldNEQ(FieldAttemptNumber, v))
}

// AttemptNumberIn applies the In predicate on the "attempt_number" field.
func AttemptNumberIn(vs ...int32) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldIn(FieldAttemptNumber, vs...))
}

// AttemptNumberNotIn applies the NotIn predicate on the "attempt_number" field.
func AttemptNumberNotIn(vs ...int32) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldNotIn(FieldAttemptNumber, vs...))
}

// AttemptNumberGT applies the GT predicate on the "attempt_number" field.
func AttemptNumberGT(v int32) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldGT(FieldAttemptNumber, v))
}

// AttemptNumberGTE applies the GTE predicate on the "attempt_number" field.
func AttemptNumberGTE(v int32) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldGTE(FieldAttemptNumber, v))
}

// AttemptNumberLT applies the LT predicate on the "attempt_number" field.
func AttemptNumberLT(v int32) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldLT(FieldAttemptNumber, v))
}

// AttemptNumberLTE applies the LTE predicate on the "attempt_number" field.
func AttemptNumberLTE(v int32) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.FieldLTE(FieldAttemptNumber, v))
}

// HasScenarioCandidate applies the HasEdge predicate on the "scenario_candidate" edge.
func HasScenarioCandidate() predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ScenarioCandidateTable, ScenarioCandidateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScenarioCandidateWith applies the HasEdge predicate on the "scenario_candidate" edge with a given conditions (other predicates).
func HasScenarioCandidateWith(preds ...predicate.ScenarioCandidate) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(func(s *sql.Selector) {
		step := newScenarioCandidateStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAnswers applies the HasEdge predicate on the "answers" edge.
func HasAnswers() predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AnswersTable, AnswersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnswersWith applies the HasEdge predicate on the "answers" edge with a given conditions (other predicates).
func HasAnswersWith(preds ...predicate.AnswerSubmission) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(func(s *sql.Selector) {
		step := newAnswersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SubmissionAttempt) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SubmissionAttempt) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SubmissionAttempt) predicate.SubmissionAttempt {
	return predicate.SubmissionAttempt(sql.NotPredicates(p))
}
