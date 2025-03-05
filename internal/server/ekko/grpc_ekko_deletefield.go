package ekko

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	ekko "ekko/api"
)

func (s *ekkoServer) DeleteField(ctx context.Context, request *ekko.DeleteFieldRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if err := s.Feature.Field.DeleteField(ctx, request); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
