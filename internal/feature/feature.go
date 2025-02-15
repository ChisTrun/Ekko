package feature

import (
	"ekko/internal/feature/field"
	"ekko/internal/repository"
	"ekko/internal/utils/extractor"
)

type Feature struct {
	Field field.Field
}

func New(repo *repository.Repository) *Feature {
	extractor := extractor.New()
	field := field.New(repo, extractor)

	return &Feature{
		Field: field,
	}
}
