package ekko

import (
	"context"
	ekko "ekko/api"
)

func (s *ekkoServer) ListAttempt(ctx context.Context, request *ekko.ListAttemptRequest) (*ekko.ListAttemptResponse, error) {
	return s.Feature.Submission.ListAttempt(ctx, request)
}
