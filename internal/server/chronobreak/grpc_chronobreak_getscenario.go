package chronobreak

import (
	"context"

	ekko "ekko/api"
)

func (s *chronobreakServer) GetScenario(ctx context.Context, request *ekko.GetScenarioRequest) (*ekko.GetScenarioResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return s.Feature.Scenario.GetScenario(ctx, request)
}
