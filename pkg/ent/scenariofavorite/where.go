// Code generated by ent, DO NOT EDIT.

package scenariofavorite

import (
	"ekko/pkg/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldEQ(FieldUserID, v))
}

// ScenarioID applies equality check predicate on the "scenario_id" field. It's identical to ScenarioIDEQ.
func ScenarioID(v uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldEQ(FieldScenarioID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldLTE(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldLTE(FieldUserID, v))
}

// ScenarioIDEQ applies the EQ predicate on the "scenario_id" field.
func ScenarioIDEQ(v uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldEQ(FieldScenarioID, v))
}

// ScenarioIDNEQ applies the NEQ predicate on the "scenario_id" field.
func ScenarioIDNEQ(v uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldNEQ(FieldScenarioID, v))
}

// ScenarioIDIn applies the In predicate on the "scenario_id" field.
func ScenarioIDIn(vs ...uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldIn(FieldScenarioID, vs...))
}

// ScenarioIDNotIn applies the NotIn predicate on the "scenario_id" field.
func ScenarioIDNotIn(vs ...uint64) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.FieldNotIn(FieldScenarioID, vs...))
}

// HasSenario applies the HasEdge predicate on the "senario" edge.
func HasSenario() predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SenarioTable, SenarioColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSenarioWith applies the HasEdge predicate on the "senario" edge with a given conditions (other predicates).
func HasSenarioWith(preds ...predicate.Scenario) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(func(s *sql.Selector) {
		step := newSenarioStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ScenarioFavorite) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ScenarioFavorite) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ScenarioFavorite) predicate.ScenarioFavorite {
	return predicate.ScenarioFavorite(sql.NotPredicates(p))
}
