package schema

import (
	ekko "ekko/api"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AnswerSubmission struct {
	ent.Schema
}

func (AnswerSubmission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (AnswerSubmission) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("submission_attempt_id"),
		field.Uint64("question_id"),
		field.Text("answer"),
		field.Float("relevance").Default(0),
		field.Float("clarity_completeness").Default(0),
		field.Float("accuracy").Default(0),
		field.Float("overall").Default(0),
		field.Int32("status").GoType(ekko.SubmissionStatus(0)),
	}
}

func (AnswerSubmission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("submission_attempt", SubmissionAttempt.Type).Ref("answers").Field("submission_attempt_id").Unique().Required(),
	}
}
