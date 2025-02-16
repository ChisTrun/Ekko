package chronobreak

import (
	"context"
	ekko "ekko/api"
)

func (s *chronobreakServer) ListScenario(ctx context.Context, request *ekko.ListScenarioRequest) (*ekko.ListScenarioResponse, error) {
	return s.Feature.Scenario.ListScenario(ctx, false, request)
}
