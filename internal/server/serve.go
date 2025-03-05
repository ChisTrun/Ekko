package server

import (
	"context"
	dbe "ekko/pkg/database/pkg/ent"
	"ekko/pkg/ent"
	"ekko/pkg/ent/migrate"
	mykit "ekko/pkg/mykit/pkg/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	pb0 "ekko/api"
	"ekko/internal/feature"
	"ekko/internal/repository"
	"ekko/internal/server/chronobreak"
	"ekko/internal/server/ekko"
	config "ekko/pkg/config"
)

// Serve ...
func Serve(cfg *config.Config) {
	service := newService(cfg, []mykit.Option{}...)
	logger := service.Logger()

	server := service.Server()

	drv, err := dbe.Open("mysql_rum", cfg.GetDatabase())
	ent := ent.NewClient(ent.Driver(drv))
	defer func() {
		if err := ent.Close(); err != nil {
			logger.Fatal("can not close ent client", zap.Error(err))
		}
	}()
	if err != nil {
		logger.Fatal("can not open ent client", zap.Error(err))
	}
	if err = ent.Schema.Create(context.Background(), migrate.WithDropIndex(true)); err != nil {
		logger.Fatal("can not init my database", zap.Error(err))
	}

	repo := repository.New(ent)

	feature := feature.New(repo)

	ekkoServer := ekko.NewServer(feature)
	chronobreakServer := chronobreak.NewServer(feature)

	grpcGatewayMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
				UseEnumNumbers:  false,
			},
		}),
	)
	service.HttpServeMux().Handle("/api/", grpcGatewayMux)

	err = pb0.RegisterEkkoHandlerServer(context.Background(), grpcGatewayMux, ekkoServer)
	if err != nil {
		logger.Fatal("can not register http sibel server", zap.Error(err))
	}

	pb0.RegisterEkkoServer(server, ekkoServer)
	pb0.RegisterChronobreakServer(server, chronobreakServer)
	// Register reflection service on gRPC server.
	// Please remove if you it's not necessary for your service
	reflection.Register(server)

	service.Serve()
}
