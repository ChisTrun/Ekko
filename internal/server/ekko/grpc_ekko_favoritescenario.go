package ekko

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	ekko "ekko/api"
)

func (s *ekkoServer) FavoriteScenario(ctx context.Context, request *ekko.FavoriteScenarioRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if err := s.Feature.Scenario.FavoriteScenario(ctx, request); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
