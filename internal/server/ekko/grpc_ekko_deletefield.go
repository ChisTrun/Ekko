package ekko

import (
	"context"
	ekko "ekko/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ekkoServer) DeleteField(ctx context.Context, request *ekko.DeleteFieldRequest) (*emptypb.Empty, error) {
	if err := s.Feature.Field.DeleteField(ctx, request); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
