package schema

import (
	ekko "ekko/api"

	"entgo.io/ent"
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
		field.Float("relevance"),
		field.Float("clarity_completeness"),
		field.Float("accuracy"),
		field.Float("overall"),
		field.Int32("status").GoType(ekko.SubmissionStatus(0)),
	}
}

func (AnswerSubmission) Edges() []ent.Edge {
	return []ent.Edge{}
}
