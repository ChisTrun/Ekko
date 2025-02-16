package feature

import (
	"ekko/internal/feature/field"
	"ekko/internal/feature/scenario"
	"ekko/internal/feature/submission"
	"ekko/internal/repository"
	"ekko/internal/utils/extractor"
)

type Feature struct {
	Field      field.Field
	Scenario   scenario.Scenario
	Submission submission.Submission
}

func New(repo *repository.Repository) *Feature {
	extractor := extractor.New()
	field := field.New(repo, extractor)
	scenario := scenario.New(repo, extractor)
	submission := submission.New(repo, extractor)

	return &Feature{
		Field:      field,
		Scenario:   scenario,
		Submission: submission,
	}
}
