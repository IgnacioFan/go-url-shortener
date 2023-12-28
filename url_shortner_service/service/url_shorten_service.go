package service

import (
	"go-url-shortener/url_shortner_service/pkg"
	"go-url-shortener/url_shortner_service/util"
)

type UrlShortenService struct {
}

func NewUrlShortenService() *UrlShortenService {
	return &UrlShortenService{}
}

func (u *UrlShortenService) CreateShortUrl(url string) (string, error) {
	// check for url
	// if the url exists in db, return
	// otherwise, create a new one
	// getIDRange from zookeeper
	zookeeper := pkg.NewZookeeperClient()
	zookeeper.GetTokenRange(1)
	// id encoded to base 62
	// add to db
	// add to redis
	return util.Encode62(1000), nil
}
