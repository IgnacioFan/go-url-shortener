package shorturl

import (
	"errors"
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
	if len(url) == 0 {
		return "", errors.New("Url is empty")
	}
	id, err := i.Repo.Create(url)
	if err != nil {
		return "", err
	}
	return base62.Encode(id), nil
}

func (i *ShortUrl) Redirect(encodedUrl string) (string, error) {
	if len(encodedUrl) > 7 {
		return "", errors.New("Short URL not found")
	}
	id, err := base62.Decode(encodedUrl)
	if err != nil {
		return "", err
	}

	originalUrl, err := i.Client.Get(encodedUrl)
	if err != nil {
		if err.Error() == "No entry" {
			return i.ReadThruCache(id, encodedUrl)
		}

		log.Fatalf("Failed to get cache entry: %v", err)
		return "", err
	} else {
		return originalUrl, err
	}
}

func (i *ShortUrl) ReadThruCache(id uint64, encodedUrl string) (string, error) {
	origanalUrl, err := i.Repo.Find(id)
	if err != nil {
		return "", err
	}
	if err = i.Client.Set(encodedUrl, origanalUrl); err != nil {
		log.Fatalf("Failed to set cache entry: %v", err)
	}
	return origanalUrl, nil
}