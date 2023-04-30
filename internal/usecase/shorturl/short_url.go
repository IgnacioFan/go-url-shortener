package shorturl

import (
	"go-url-shortener/internal/entity"
	"go-url-shortener/internal/usecase/base62"
	"go-url-shortener/pkg/redis"
	"log"
)

type ShortUrl struct {
	Client redis.RedisClient
	Repo   entity.ShortUrlRepository
}

func NewShortUrlUsecase(repo entity.ShortUrlRepository, client redis.RedisClient) entity.ShortUrlUsecase {
	return &ShortUrl{
		Client: client,
		Repo:   repo,
	}
}

func (i *ShortUrl) Create(url string) (string, error) {
	id, err := i.Repo.Create(url)
	if err != nil {
		return "", err
	}
	return base62.Encode(id), nil
}

func (i *ShortUrl) Redirect(code string) (string, error) {
	id, err := base62.Decode(code)
	if err != nil {
		return "", err
	}

	url, err := i.Client.Get(code)
	if err != nil && err.Error() == "No entry" {
		url, err := i.Repo.Find(id)
		_ = i.Client.Set(code, url)

		return url, err
	} else if err != nil {
		log.Fatalf("Failed to get cache entry: %v", err)
		return "", err
	} else {
		return url, err
	}
}

func (i *ShortUrl) Delete(code string) error {
	id, err := base62.Decode(code)
	if err != nil {
		return err
	}
	if err = i.Repo.Delete(id); err != nil {
		return err
	}
	if err = i.Client.Del(code); err != nil {
		return err
	}
	return nil
}
