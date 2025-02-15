package ekko

import (
	"context"
	ekko "ekko/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ekkoServer) UpdateField(ctx context.Context, request *ekko.UpdateFieldRequest) (*emptypb.Empty, error) {
	if err := s.Feature.Field.UpdateField(ctx, request); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
