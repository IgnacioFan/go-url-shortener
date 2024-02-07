package url_repo

import (
	"go-url-shortener/internal/adpater/postgres"
	"go-url-shortener/internal/entity"
)

type UrlRepository interface {
	Create(longURL, shortUrl string) error
	FindBy(shortUrl string) (string, error)
}

type Url struct {
	postgres.Postgres
}

func NewShortUrlRepo(postgres *postgres.Postgres) UrlRepository {
	return &Url{*postgres}
}

func (i *Url) Create(longURL, shortUrl string) error {
	res := &entity.Url{LongURL: longURL, ShortURL: shortUrl}
	if err := i.DB.Create(res).Error; err != nil {
		return err
	}
	return nil
}

func (i *Url) FindBy(shortUrl string) (string, error) {
	res := &entity.Url{}
	if err := i.DB.Where("short_url = ?", shortUrl).First(res).Error; err != nil {
		return "", err
	}
	return res.LongURL, nil
}
