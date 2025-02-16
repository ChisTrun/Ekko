package chronobreak

import (
	"context"
	ekko "ekko/api"
)

func (s *chronobreakServer) GetScenario(ctx context.Context, request *ekko.GetScenarioRequest) (*ekko.GetScenarioResponse, error) {
	return s.Feature.Scenario.GetScenario(ctx, request)
}
