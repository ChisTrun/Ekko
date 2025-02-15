package repository

import (
	"ekko/internal/repository/field"
	"ekko/internal/repository/question"
	"ekko/internal/repository/scenario"
	"ekko/package/ent"
)

type Repository struct {
	Field    field.Field
	Question question.Question
	Scenario scenario.Scenario
	Ent      *ent.Client
}

func New(ent *ent.Client) *Repository {
	field := field.New(ent)
	question := question.New()
	scenario := scenario.New(ent, question)

	return &Repository{
		Field:    field,
		Question: question,
		Scenario: scenario,
		Ent:      ent,
	}
}
