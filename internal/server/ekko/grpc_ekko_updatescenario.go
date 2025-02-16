package ekko

import (
	"context"
	ekko "ekko/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ekkoServer) UpdateScenario(ctx context.Context, request *ekko.UpdateScenarioRequest) (*emptypb.Empty, error) {
	if err := s.Feature.Scenario.UpdateScenario(ctx, request); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
