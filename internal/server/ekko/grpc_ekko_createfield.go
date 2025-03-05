package ekko

import (
	"context"

	ekko "ekko/api"
)

func (s *ekkoServer) CreateField(ctx context.Context, request *ekko.CreateFieldRequest) (*ekko.CreateFieldResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return s.Feature.Field.CreateField(ctx, request)
}
