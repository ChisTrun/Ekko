package ekko

import (
	"context"
	ekko "ekko/api"
)

func (s *ekkoServer) CreateScenario(ctx context.Context, request *ekko.CreateScenarioRequest) (*ekko.CreateScenarioResponse, error) {
	return s.Feature.Scenario.CreateScenario(ctx, request)
}
