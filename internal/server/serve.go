package server

import (
	"context"
	dbe "ekko/pkg/database/pkg/ent"
	"ekko/pkg/ent"
	"ekko/pkg/ent/migrate"
	mykit "ekko/pkg/mykit/pkg/api"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	pb0 "ekko/api"
	"ekko/internal/bulbasaur"
	"ekko/internal/feature"
	"ekko/internal/rabbit"
	"ekko/internal/repository"
	"ekko/internal/server/chronobreak"
	"ekko/internal/server/ekko"
	"ekko/internal/utils/redis"
	config "ekko/pkg/config"
)

func customMetadataAnnotator(ctx context.Context, req *http.Request) metadata.MD {
	md := metadata.MD{}

	// Map tất cả các header có prefix "x-" vào metadata
	for name, values := range req.Header {
		lowerName := strings.ToLower(name)
		if strings.HasPrefix(lowerName, "x-") {
			md.Append(lowerName, values...)
		}
	}

	return md
}

// Serve ...
func Serve(cfg *config.Config) {
	service := newService(cfg, []mykit.Option{}...)
	logger := service.Logger()

	server := service.Server()

	drv, err := dbe.Open("mysql_ekko", cfg.GetDatabase())
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

	rabbitMQ := rabbit.New(cfg.Rabbitmq)

	bulbasaur := bulbasaur.New(cfg.Bulbasaur)

	repo := repository.New(ent, rabbitMQ, bulbasaur)

	go rabbitMQ.Consume(context.Background(), repo.Submission.ReceiveResponse)

	feature := feature.New(repo)

	redis := redis.New(cfg.Flags.EnableRedis, cfg)

	ekkoServer := ekko.NewServer(feature)
	chronobreakServer := chronobreak.NewServer(feature, redis)

	grpcGatewayMux := runtime.NewServeMux(
		runtime.WithMetadata(customMetadataAnnotator),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
				UseEnumNumbers:  false,
			},
		}),
	)
	service.HttpServeMux().Handle("/ekko/", grpcGatewayMux)
	service.HttpServeMux().Handle("/chronobreak/", grpcGatewayMux)

	err = pb0.RegisterEkkoHandlerServer(context.Background(), grpcGatewayMux, ekkoServer)
	if err != nil {
		logger.Fatal("can not register http sibel server", zap.Error(err))
	}
	err = pb0.RegisterChronobreakHandlerServer(context.Background(), grpcGatewayMux, chronobreakServer)
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
