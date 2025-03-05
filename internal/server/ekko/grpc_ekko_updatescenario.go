package ekko

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	ekko "ekko/api"
)

func (s *ekkoServer) UpdateScenario(ctx context.Context, request *ekko.UpdateScenarioRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if err := s.Feature.Scenario.UpdateScenario(ctx, request); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
