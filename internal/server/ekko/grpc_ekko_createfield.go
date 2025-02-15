package ekko

import (
	"context"
	ekko "ekko/api"
)

func (s *ekkoServer) CreateField(ctx context.Context, request *ekko.CreateFieldRequest) (*ekko.CreateFieldResponse, error) {
	return s.Feature.Field.CreateField(ctx, request)
}
