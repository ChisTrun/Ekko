package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ScenarioCandidate struct {
	ent.Schema
}

func (ScenarioCandidate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (ScenarioCandidate) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("candidate_id"),
		field.Uint64("scenario_id"),
	}
}

func (ScenarioCandidate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("scenario", Scenario.Type).Ref("candidates").Field("scenario_id").Unique().Required(),
		edge.To("attempts", SubmissionAttempt.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
