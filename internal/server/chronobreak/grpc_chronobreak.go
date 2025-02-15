package chronobreak

import (
	ekko "ekko/api"
	"ekko/internal/feature"
)

func NewServer(feature *feature.Feature) ekko.ChronobreakServer {
	return &chronobreakServer{
		Feature: feature,
	}
}

type chronobreakServer struct {
	ekko.UnimplementedChronobreakServer
	Feature *feature.Feature
}
