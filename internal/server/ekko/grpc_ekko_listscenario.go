package ekko

import (
	"context"
	ekko "ekko/api"
)

func (s *ekkoServer) ListScenario(ctx context.Context, request *ekko.ListScenarioRequest) (*ekko.ListScenarioResponse, error) {
	return s.Feature.Scenario.ListScenario(ctx, true, request)
}
