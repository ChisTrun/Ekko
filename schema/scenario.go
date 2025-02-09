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
	}
}

func (Scenario) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("candidates", ScenarioCandidate.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
