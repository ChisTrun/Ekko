package redis

import (
	"context"
	config "ekko/pkg/config"
	"fmt"
	"time"

	re "github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Redis interface {
	Set(ctx context.Context, key string, value proto.Message, expireTime time.Duration) (bool, error)
	Get(ctx context.Context, key string) ([]byte, error)
}

type redis struct {
	redis     *re.Client
	namespace string
}

func New(enable bool, cfg *config.Config) Redis {
	if !enable {
		return Noop()
	}

	return &redis{
		redis: re.NewClient(&re.Options{
			Addr: cfg.Redis.Address,
		}),
		namespace: cfg.Redis.Namespace, // Assuming namespace is part of the config
	}
}

func (r *redis) withNamespace(key string) string {
	return fmt.Sprintf("%s:%s", r.namespace, key)
}

func (r *redis) Set(ctx context.Context, key string, value proto.Message, expireTime time.Duration) (bool, error) {
	namespacedKey := r.withNamespace(key)
	jsonData, err := protojson.Marshal(value)
	if err != nil {
		return false, err
	}
	return r.redis.Set(ctx, namespacedKey, string(jsonData), expireTime).Err() == nil, nil
}

func (r *redis) Get(ctx context.Context, key string) ([]byte, error) {
	namespacedKey := r.withNamespace(key)
	val, err := r.redis.Get(ctx, namespacedKey).Result()
	if err == re.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return []byte(val), nil
}
