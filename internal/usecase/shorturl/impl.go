package shorturl

import (
	"errors"
	"go-url-shortener/internal/repository/postgres"
	"go-url-shortener/internal/repository/redis"
	"go-url-shortener/internal/usecase"
	"go-url-shortener/internal/usecase/base62"
	"log"
)

type Impl struct {
	Cache redis.UrlCache
	Repo  postgres.UrlRepository
}

func NewShortUrl(cache redis.UrlCache, repo postgres.UrlRepository) usecase.ShortUrl {
	return &Impl{
		Cache: cache,
		Repo:  repo,
	}
}

func (i *Impl) Create(url string) (string, error) {
	if len(url) == 0 {
		return "", errors.New("Url is empty")
	}
	id, err := i.Repo.Create(url)
	if err != nil {
		return "", err
	}
	return base62.Encode(id), nil
}

func (i *Impl) Redirect(encodedUrl string) (string, error) {
	if len(encodedUrl) > 7 {
		return "", errors.New("Short URL not found")
	}
	id, err := base62.Decode(encodedUrl)
	if err != nil {
		return "", err
	}

	originalUrl, err := i.Cache.Get(encodedUrl)
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

func (i *Impl) ReadThruCache(id uint64, encodedUrl string) (string, error) {
	origanalUrl, err := i.Repo.Find(id)
	if err != nil {
		return "", err
	}
	if err = i.Cache.Set(encodedUrl, origanalUrl); err != nil {
		log.Fatalf("Failed to set cache entry: %v", err)
	}
	return origanalUrl, nil
}
