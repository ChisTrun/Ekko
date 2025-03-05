package ekko

import (
	"context"

	ekko "ekko/api"
)

func (s *ekkoServer) ListAllSubmission(ctx context.Context, request *ekko.ListAllSubmissionRequest) (*ekko.ListAllSubmissionResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return s.Feature.Submission.ListAllSubmission(ctx, request)
}
