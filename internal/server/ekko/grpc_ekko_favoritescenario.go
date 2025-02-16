package ekko

import (
	"context"
	ekko "ekko/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ekkoServer) FavoriteScenario(ctx context.Context, request *ekko.FavoriteScenarioRequest) (*emptypb.Empty, error) {
	if err := s.Feature.Scenario.FavoriteScenario(ctx, request); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
