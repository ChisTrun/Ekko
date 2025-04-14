package repository

import (
	"ekko/internal/rabbit"
	"ekko/internal/repository/field"
	"ekko/internal/repository/question"
	"ekko/internal/repository/scenario"
	"ekko/internal/repository/submission"
	"ekko/pkg/ent"
)

type Repository struct {
	Field      field.Field
	Question   question.Question
	Scenario   scenario.Scenario
	Submission submission.Submission
	Ent        *ent.Client
}

func New(ent *ent.Client, rabbitMQ rabbit.Rabbit) *Repository {
	field := field.New(ent)
	question := question.New()
	scenario := scenario.New(ent, question)
	submission := submission.New(ent, rabbitMQ)

	return &Repository{
		Field:      field,
		Question:   question,
		Scenario:   scenario,
		Ent:        ent,
		Submission: submission,
	}
}
