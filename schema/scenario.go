package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Scenario struct {
	ent.Schema
}

func (Scenario) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (Scenario) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("bm_id"),
		field.String("name"),
		field.Text("description"),
		field.Float("rating").Default(0),
		field.Int32("total_rating").Default(0),
		field.Int32("participants").Default(0),
	}
}

func (Scenario) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("candidates", ScenarioCandidate.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("favorites", ScenarioFavorite.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("field", ScenarioField.Type).Ref("senarios"),
	}
}
