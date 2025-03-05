package ekko

import (
	"context"

	ekko "ekko/api"
)

func (s *ekkoServer) ListAttempt(ctx context.Context, request *ekko.ListAttemptRequest) (*ekko.ListAttemptResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return s.Feature.Submission.ListAttempt(ctx, request)
}
