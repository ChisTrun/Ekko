package ekko

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	ekko "ekko/api"
)

func (s *ekkoServer) DeleteScenario(ctx context.Context, request *ekko.DeleteScenarioRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if err := s.Feature.Scenario.DeleteScenario(ctx, request); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
