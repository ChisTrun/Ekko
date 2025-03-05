package ekko

import (
	"context"

	"ekko/api"
)

func (s *ekkoServer) TestApi(ctx context.Context, request *ekko.TestRequest) (*ekko.TestResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &ekko.TestResponse{}, nil
}
