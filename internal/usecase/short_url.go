package usecase

import (
	"errors"
	"fmt"
	"go-url-shortener/internal/repository/postgres"
)

type ShortUrlUsecase interface {
	Create(url string) (string, error)
	Redirect(url string) (string, error)
}

type ShortUrl struct {
	Repo postgres.UrlRepository
}

func NewShortUrl(repo postgres.UrlRepository) *ShortUrl {
	return &ShortUrl{Repo: repo}
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

func (s *ShortUrl) Redirect(url string) (string, error) {
	if len(url) > 7 {
		return "", errors.New("Short URL not found")
	}
	id, err := Decode(url)
	if err != nil {
		return "", err
	}
	fmt.Println(id)
	return "https://example.com/foobar", nil
}
