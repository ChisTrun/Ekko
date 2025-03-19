package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ScenarioField struct {
	ent.Schema
}

func (ScenarioField) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (ScenarioField) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (ScenarioField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("senarios", Scenario.Type),
	}
}
