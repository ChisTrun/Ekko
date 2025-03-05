package ekko

import (
	"context"

	ekko "ekko/api"
)

func (s *ekkoServer) GetAttempt(ctx context.Context, request *ekko.GetAttemptRequest) (*ekko.GetAttemptResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return s.Feature.Submission.GetAttempt(ctx, request)
}
