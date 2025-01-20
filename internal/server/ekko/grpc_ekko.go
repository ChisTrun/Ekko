package ekko

import (
	ekko "ekko/api"
)

func NewServer() ekko.EkkoServer {
	return &ekkoServer{}
}

type ekkoServer struct {
	ekko.UnimplementedEkkoServer
}
