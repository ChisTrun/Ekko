package scenario

import "ekko/internal/repository"

type Scenario interface {
}

type scenario struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) Scenario {
	return &scenario{
		repo: repo,
	}
}


