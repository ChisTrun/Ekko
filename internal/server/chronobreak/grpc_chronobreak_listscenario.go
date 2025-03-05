package chronobreak

import (
	"context"

	ekko "ekko/api"
)

func (s *chronobreakServer) ListScenario(ctx context.Context, request *ekko.ListScenarioRequest) (*ekko.ListScenarioResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return s.Feature.Scenario.ListScenario(ctx, false, request)
}
