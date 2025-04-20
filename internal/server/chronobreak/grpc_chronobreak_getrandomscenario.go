package chronobreak

import (
	"context"
	ekko "ekko/api"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *chronobreakServer) GetRandomScenario(ctx context.Context, request *emptypb.Empty) (*ekko.GetScenarioResponse, error) {
	data, err := s.Redis.Get(ctx, "random_scenario")
	if err != nil {
		return nil, err
	}
	if data != nil {
		mss := &ekko.GetScenarioResponse{}
		if err := protojson.Unmarshal(data, mss); err != nil {
			return nil, err
		}
		return mss, nil
	}

	mss, err := s.Feature.Scenario.GetRandomScenario(ctx)
	if err != nil {
		return nil, err
	}

	go func(ctx context.Context, mss proto.Message) {
		if _, err := s.Redis.Set(ctx, "random_scenario", mss, 48*time.Hour); err != nil {
		}
	}(context.Background(), mss)

	return mss, nil
}
