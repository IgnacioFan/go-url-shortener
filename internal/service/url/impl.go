package url

import (
	"go-url-shortener/internal/service/base62"
	"go-url-shortener/pkg/zookeeper"
)

type UrlService interface {
	GenerateShortURL(longUrl string) (string, error)
	OriginalURL(shortUrl string) (string, error)
}

type Impl struct {
	ZkClient *zookeeper.Impl
}

func InitUrl(zkClient *zookeeper.Impl) *Impl {
	return &Impl{ZkClient: zkClient}
}

func (i *Impl) GenerateShortURL(longURL string) (string, error) {
	uniqueId := i.ZkClient.GetCounter()
	shortUrl := base62.Encode(uint64(uniqueId))
	// TODO: update the counter
	// TODO: store it into the database
	return shortUrl, nil
}

func (i *Impl) OriginalURL(shortURL string) (string, error)  {
	return "https://example.com/foobar", nil
}
