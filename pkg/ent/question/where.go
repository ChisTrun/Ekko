// Code generated by ent, DO NOT EDIT.

package question

import (
	"ekko/pkg/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldUpdatedAt, v))
}

// ScenarioID applies equality check predicate on the "scenario_id" field. It's identical to ScenarioIDEQ.
func ScenarioID(v uint64) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldScenarioID, v))
}

// Criteria applies equality check predicate on the "criteria" field. It's identical to CriteriaEQ.
func Criteria(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldCriteria, v))
}

// Hint applies equality check predicate on the "hint" field. It's identical to HintEQ.
func Hint(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldHint, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldContent, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldUpdatedAt, v))
}

// ScenarioIDEQ applies the EQ predicate on the "scenario_id" field.
func ScenarioIDEQ(v uint64) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldScenarioID, v))
}

// ScenarioIDNEQ applies the NEQ predicate on the "scenario_id" field.
func ScenarioIDNEQ(v uint64) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldScenarioID, v))
}

// ScenarioIDIn applies the In predicate on the "scenario_id" field.
func ScenarioIDIn(vs ...uint64) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldScenarioID, vs...))
}

// ScenarioIDNotIn applies the NotIn predicate on the "scenario_id" field.
func ScenarioIDNotIn(vs ...uint64) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldScenarioID, vs...))
}

// CriteriaEQ applies the EQ predicate on the "criteria" field.
func CriteriaEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldCriteria, v))
}

// CriteriaNEQ applies the NEQ predicate on the "criteria" field.
func CriteriaNEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldCriteria, v))
}

// CriteriaIn applies the In predicate on the "criteria" field.
func CriteriaIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldCriteria, vs...))
}

// CriteriaNotIn applies the NotIn predicate on the "criteria" field.
func CriteriaNotIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldCriteria, vs...))
}

// CriteriaGT applies the GT predicate on the "criteria" field.
func CriteriaGT(v string) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldCriteria, v))
}

// CriteriaGTE applies the GTE predicate on the "criteria" field.
func CriteriaGTE(v string) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldCriteria, v))
}

// CriteriaLT applies the LT predicate on the "criteria" field.
func CriteriaLT(v string) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldCriteria, v))
}

// CriteriaLTE applies the LTE predicate on the "criteria" field.
func CriteriaLTE(v string) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldCriteria, v))
}

// CriteriaContains applies the Contains predicate on the "criteria" field.
func CriteriaContains(v string) predicate.Question {
	return predicate.Question(sql.FieldContains(FieldCriteria, v))
}

// CriteriaHasPrefix applies the HasPrefix predicate on the "criteria" field.
func CriteriaHasPrefix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasPrefix(FieldCriteria, v))
}

// CriteriaHasSuffix applies the HasSuffix predicate on the "criteria" field.
func CriteriaHasSuffix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasSuffix(FieldCriteria, v))
}

// CriteriaEqualFold applies the EqualFold predicate on the "criteria" field.
func CriteriaEqualFold(v string) predicate.Question {
	return predicate.Question(sql.FieldEqualFold(FieldCriteria, v))
}

// CriteriaContainsFold applies the ContainsFold predicate on the "criteria" field.
func CriteriaContainsFold(v string) predicate.Question {
	return predicate.Question(sql.FieldContainsFold(FieldCriteria, v))
}

// HintEQ applies the EQ predicate on the "hint" field.
func HintEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldHint, v))
}

// HintNEQ applies the NEQ predicate on the "hint" field.
func HintNEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldHint, v))
}

// HintIn applies the In predicate on the "hint" field.
func HintIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldHint, vs...))
}

// HintNotIn applies the NotIn predicate on the "hint" field.
func HintNotIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldHint, vs...))
}

// HintGT applies the GT predicate on the "hint" field.
func HintGT(v string) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldHint, v))
}

// HintGTE applies the GTE predicate on the "hint" field.
func HintGTE(v string) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldHint, v))
}

// HintLT applies the LT predicate on the "hint" field.
func HintLT(v string) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldHint, v))
}

// HintLTE applies the LTE predicate on the "hint" field.
func HintLTE(v string) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldHint, v))
}

// HintContains applies the Contains predicate on the "hint" field.
func HintContains(v string) predicate.Question {
	return predicate.Question(sql.FieldContains(FieldHint, v))
}

// HintHasPrefix applies the HasPrefix predicate on the "hint" field.
func HintHasPrefix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasPrefix(FieldHint, v))
}

// HintHasSuffix applies the HasSuffix predicate on the "hint" field.
func HintHasSuffix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasSuffix(FieldHint, v))
}

// HintEqualFold applies the EqualFold predicate on the "hint" field.
func HintEqualFold(v string) predicate.Question {
	return predicate.Question(sql.FieldEqualFold(FieldHint, v))
}

// HintContainsFold applies the ContainsFold predicate on the "hint" field.
func HintContainsFold(v string) predicate.Question {
	return predicate.Question(sql.FieldContainsFold(FieldHint, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Question {
	return predicate.Question(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Question {
	return predicate.Question(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Question {
	return predicate.Question(sql.FieldContainsFold(FieldContent, v))
}

// HasScenario applies the HasEdge predicate on the "scenario" edge.
func HasScenario() predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ScenarioTable, ScenarioColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScenarioWith applies the HasEdge predicate on the "scenario" edge with a given conditions (other predicates).
func HasScenarioWith(preds ...predicate.Scenario) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := newScenarioStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAnswers applies the HasEdge predicate on the "answers" edge.
func HasAnswers() predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AnswersTable, AnswersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnswersWith applies the HasEdge predicate on the "answers" edge with a given conditions (other predicates).
func HasAnswersWith(preds ...predicate.AnswerSubmission) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := newAnswersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Question) predicate.Question {
	return predicate.Question(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Question) predicate.Question {
	return predicate.Question(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Question) predicate.Question {
	return predicate.Question(sql.NotPredicates(p))
}
