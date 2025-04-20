package chronobreak

import (
	ekko "ekko/api"
	"ekko/internal/feature"
	"ekko/internal/utils/redis"
)

func NewServer(feature *feature.Feature, redis redis.Redis) ekko.ChronobreakServer {
	return &chronobreakServer{
		Feature: feature,
		Redis:   redis,
	}
}

type chronobreakServer struct {
	ekko.UnimplementedChronobreakServer
	Feature *feature.Feature
	Redis   redis.Redis
}
