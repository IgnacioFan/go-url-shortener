package redis

import (
	"context"
	"fmt"
	"go-url-shortener/deployment/config"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

func InitClient() (*redis.Client, error) {
	config, err := config.New()
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password:     config.Redis.Password,
		DB:           config.Redis.DB,
		PoolSize:     config.Redis.MaxPoolSize,
		MinIdleConns: config.Redis.MinIdleConns,
	})
	ctx, cancel := context.WithTimeout(context.Background(), config.Redis.DialTimeout)
	defer cancel()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.Wrap(err, "redis error:")
	}

	return client, nil
}
