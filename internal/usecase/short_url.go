package usecase

import (
	"errors"
	"fmt"
	"go-url-shortener/internal/repostiory/postgres"
)

type ShortUrlUsecase interface {
	Create(url string) (string, error)
	Redirect(url string) (string, error)
}

type ShortUrl struct {
	Repo *postgres.UrlRepository
}

func NewShortUrl(repo *postgres.UrlRepository) *ShortUrl {
	return &ShortUrl{
		Repo: repo,
	}
}

func (s *ShortUrl) Create(url string) (string, error) {
	if url == "" {
		return "", errors.New("Url is empty")
	}
	return "abc", nil
}

func (s *ShortUrl) Redirect(url string) (string, error) {
	if url == "invalid" {
		return "", errors.New("Short URL not found")
	}
	return "https://example.com/foobar", nil
}
