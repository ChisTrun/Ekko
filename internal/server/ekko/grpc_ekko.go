package ekko

import (
	ekko "ekko/api"
	"ekko/internal/feature"
)

func NewServer(feature *feature.Feature) ekko.EkkoServer {
	return &ekkoServer{
		Feature: feature,
	}
}

type ekkoServer struct {
	ekko.UnimplementedEkkoServer
	Feature *feature.Feature
}
