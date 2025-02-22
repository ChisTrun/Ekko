package server

import (
	"context"
	pb0 "ekko/api"
	"ekko/internal/feature"
	"ekko/internal/repository"
	"ekko/internal/server/chronobreak"
	"ekko/internal/server/ekko"
	"ekko/internal/utils/logging"
	"ekko/package/config"
	"ekko/package/ent"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func Serve(cfg *config.Config) {
	client, err := ent.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	repo := repository.New(client)

	feature := feature.New(repo)

	pb0.RegisterChronobreakServer(grpcServer, chronobreak.NewServer(feature))
	pb0.RegisterEkkoServer(grpcServer, ekko.NewServer(feature))

	logging.Logger(context.Background()).Info(fmt.Sprintf("server is runing on: %v:%v", cfg.Database.Username, cfg.Server.Port))
	grpcServer.Serve(lis)
}
