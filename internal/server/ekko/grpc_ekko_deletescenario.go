package ekko

import (
	"context"
	ekko "ekko/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ekkoServer) DeleteScenario(ctx context.Context, request *ekko.DeleteScenarioRequest) (*emptypb.Empty, error) {
	return nil, nil
}
