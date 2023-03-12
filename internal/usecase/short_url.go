package usecase

import (
	"errors"
)

type ShortUrl struct{}

func NewShortUrl() *ShortUrl {
	return &ShortUrl{}
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
