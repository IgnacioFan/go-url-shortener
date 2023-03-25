package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

type UrlCache interface {
	Get(key string) (string, error)
	Set(key, val string) error
}

type Url struct {
	Client *redis.Client
}

func (u *Url) Get(key string) (string, error) {
	if str, err := u.Client.Get(key).Result(); err == redis.Nil {
		return "", errors.New("No entry")
	} else {
		return str, err
	}
}

func (u *Url) Set(key, val string) error {
	expiration := time.Duration(24 * time.Hour)
	if _, err := u.Client.Set(key, val, expiration).Result(); err != nil {
		return err
	}
	return nil
}
