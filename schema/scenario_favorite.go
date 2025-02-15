package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ScenarioFavorite struct {
	ent.Schema
}

func (ScenarioFavorite) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("user_id"),
		field.Uint64("scenario_id"),
	}
}

func (ScenarioFavorite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("senario", Scenario.Type).Ref("favorites").Field("scenario_id").Unique().Required(),
	}
}

func (ScenarioFavorite) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
