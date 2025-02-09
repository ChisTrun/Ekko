package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Question struct {
	ent.Schema
}

func (Question) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("sentence_id"),
		field.Text("criteria"),
		field.Text("hint"),
		field.Text("content"),
	}
}

func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("scenario", Scenario.Type).
			Ref("questions").
			Field("sentence_id").
			Required().
			Unique(),
	}
}
