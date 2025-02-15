package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SubmissionAttempt struct {
	ent.Schema
}

func (SubmissionAttempt) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (SubmissionAttempt) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("scenario_candidate_id"),
		field.Int32("attempt_number"),
	}
}

func (SubmissionAttempt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("scenario_candidate", ScenarioCandidate.Type).Ref("attempts").Field("scenario_candidate_id").Unique().Required(),
		edge.To("answers", AnswerSubmission.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (SubmissionAttempt) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("scenario_candidate_id", "attempt_number").Unique(),
	}
}
