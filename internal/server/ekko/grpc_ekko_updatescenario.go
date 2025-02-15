package ekko

import (
	"context"
	ekko "ekko/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ekkoServer) UpdateScenario(ctx context.Context, request *ekko.UpdateScenarioRequest) (*emptypb.Empty, error) {
	return nil, nil
}
