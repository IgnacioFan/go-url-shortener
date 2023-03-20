package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type UrlRepository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, val string) error
}

type Url struct {
	Client *redis.Client
}

func (u *Url) Get(ctx context.Context, key string) (string, error) {
	originalUrl, err := u.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}

func (u *Url) Set(ctx context.Context, key, val string) error {
	expiration := time.Duration(24 * time.Hour)
	if _, err := u.Client.Set(ctx, key, val, expiration).Result(); err != nil {
		return err
	}
	return nil
}
