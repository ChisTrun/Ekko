package ekko

import (
	"context"

	ekko "ekko/api"
)

func (s *ekkoServer) ListScenario(ctx context.Context, request *ekko.ListScenarioRequest) (*ekko.ListScenarioResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return s.Feature.Scenario.ListScenario(ctx, true, request)
}
