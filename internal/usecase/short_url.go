package usecase

import (
	"context"
	"errors"
	"go-url-shortener/internal/repository/postgres"
	"go-url-shortener/internal/repository/redis"
	"log"
)

type ShortUrlUsecase interface {
	Create(url string) (string, error)
	Redirect(url string) (string, error)
}

type ShortUrl struct {
	Client  redis.UrlCache
	Context context.Context
	Repo    postgres.UrlRepository
}

func NewShortUrl(client redis.UrlCache, ctx context.Context, repo postgres.UrlRepository) *ShortUrl {
	return &ShortUrl{
		Client:  client,
		Context: ctx,
		Repo:    repo,
	}
}

func (s *ShortUrl) Create(url string) (string, error) {
	if len(url) == 0 {
		return "", errors.New("Url is empty")
	}
	id, err := s.Repo.Create(url)
	if err != nil {
		return "", err
	}
	return Encode(id), nil
}

func (s *ShortUrl) Redirect(encodedUrl string) (string, error) {
	if len(encodedUrl) > 7 {
		return "", errors.New("Short URL not found")
	}
	id, err := Decode(encodedUrl)
	if err != nil {
		return "", err
	}

	originalUrl, err := s.Client.Get(s.Context, encodedUrl)
	if err != nil {
		if err.Error() == "No entry" {
			return s.ReadThruCache(id, encodedUrl)
		}

		log.Fatalf("Failed to get cache entry: %v", err)
		return "", err
	} else {
		return originalUrl, err
	}
}

func (s *ShortUrl) ReadThruCache(id uint64, encodedUrl string) (string, error) {
	origanalUrl, err := s.Repo.Find(id)
	if err != nil {
		return "", err
	}
	if err = s.Client.Set(s.Context, encodedUrl, origanalUrl); err != nil {
		log.Fatalf("Failed to set cache entry: %v", err)
	}
	return origanalUrl, nil
}
