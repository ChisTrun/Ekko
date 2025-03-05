package ekko

import (
	"context"

	ekko "ekko/api"
)

func (s *ekkoServer) CreateScenario(ctx context.Context, request *ekko.CreateScenarioRequest) (*ekko.CreateScenarioResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return s.Feature.Scenario.CreateScenario(ctx, request)
}
