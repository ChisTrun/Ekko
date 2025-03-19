package ekko

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	ekko "ekko/api"
)

func (s *ekkoServer) RatingScenario(ctx context.Context, request *ekko.RatingScenarioRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if err := s.Feature.Scenario.RatingScenario(ctx, request); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
