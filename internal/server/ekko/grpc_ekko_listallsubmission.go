package ekko

import (
	"context"
	ekko "ekko/api"
)

func (s *ekkoServer) ListAllSubmission(ctx context.Context, request *ekko.ListAllSubmissionRequest) (*ekko.ListAllSubmissionResponse, error) {
	return s.Feature.Submission.ListAllSubmission(ctx, request)
}
