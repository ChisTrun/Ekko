package ekko

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	ekko "ekko/api"
)

func (s *ekkoServer) UpdateField(ctx context.Context, request *ekko.UpdateFieldRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if err := s.Feature.Field.UpdateField(ctx, request); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
