package chronobreak

import (
	"context"

	ekko "ekko/api"
)

func (s *chronobreakServer) ListField(ctx context.Context, request *ekko.ListFieldRequest) (*ekko.ListFieldResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return s.Feature.Field.ListField(ctx, request)
}
