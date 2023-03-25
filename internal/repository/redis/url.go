package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type UrlCache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, val string) error
}

type Url struct {
	Client *redis.Client
}

func (u *Url) Get(ctx context.Context, key string) (string, error) {
	if str, err := u.Client.Get(ctx, key).Result(); err == redis.Nil {
		return "", errors.New("No entry")
	} else {
		return str, err
	}
}

func (u *Url) Set(ctx context.Context, key, val string) error {
	expiration := time.Duration(24 * time.Hour)
	if _, err := u.Client.Set(ctx, key, val, expiration).Result(); err != nil {
		return err
	}
	return nil
}
