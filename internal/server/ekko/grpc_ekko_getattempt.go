package ekko

import (
	"context"
	ekko "ekko/api"
)

func (s *ekkoServer) GetAttempt(ctx context.Context, request *ekko.GetAttemptRequest) (*ekko.GetAttemptResponse, error) {
	return s.Feature.Submission.GetAttempt(ctx, request)
}
