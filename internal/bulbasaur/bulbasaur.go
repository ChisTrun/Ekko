package bulbasaur

import (
	"context"
	bul "ekko/pkg/bulbasaur/api"
	"ekko/pkg/logger/pkg/logging"
	"flag"
	"fmt"

	cfg "ekko/pkg/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Bulbasaur interface {
	FindUserByNameRequest(ctx context.Context, content string) []uint64
}

type bulbasaur struct {
	client bul.VenusaurClient
}

func New(cfg *cfg.InternalService) Bulbasaur {
	ctx := context.Background()
	if cfg == nil {
		logging.Logger(ctx).Error("bulbasaur config is nil")
		return &Noop{}
	}

	serverAddr := flag.String("addr", fmt.Sprintf("%v:%v", cfg.Address, cfg.Port), "The server address in the format of host:port")
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		logging.Logger(ctx).Error(fmt.Sprintf("failed to connect: %v", err.Error()))
		return &Noop{}
	}

	return &bulbasaur{
		client: bul.NewVenusaurClient(conn),
	}
}

func (b *bulbasaur) FindUserByNameRequest(ctx context.Context, content string) []uint64 {
	rsp, err := b.client.FindUserByName(ctx, &bul.FindUserByNameRequest{
		Name: content,
	})
	if err != nil {
		logging.Logger(ctx).Error(fmt.Sprintf("failed to find user by name: %v", err.Error()))
		return []uint64{}
	}

	return rsp.Ids
}
