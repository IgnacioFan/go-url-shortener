package usecase

import (
	"errors"
)

type ShortUrl struct{}

func New() *ShortUrl {
	return &ShortUrl{}
}

func (s *ShortUrl) Create(url string) (string, error) {
	if url == "" {
		return "", errors.New("Url is empty")
	}
	return "abc", nil
}

func (s *ShortUrl) Redirect(url string) (string, error) {
	return "http://www.example.com/abc", nil
}
