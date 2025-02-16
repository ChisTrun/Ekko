package ekko

import (
	"context"
	ekko "ekko/api"
)

func (s *ekkoServer) SubmitAnswer(ctx context.Context, request *ekko.SubmitAnswerRequest) (*ekko.SubmitAnswerResponse, error) {
	return s.Feature.Submission.SubmitAnswer(ctx, request)
}
