package service

import "fmt"

type UrlShortenService struct {
}

func NewUrlShortenService() *UrlShortenService {
	return &UrlShortenService{}
}

func (u *UrlShortenService) CreateShortUrl(url string) (string, error) {
	fmt.Println("received a long URL and processing", url)
	return "", nil
}
