package repository

import (
	"errors"
	"go-url-shortener/internal/entity"
	"go-url-shortener/pkg/postgres"

	"gorm.io/gorm"
)

type ShortUrl struct {
	postgres.Postgres
}

func NewShortUrlRepo(postgres *postgres.Postgres) entity.ShortUrlRepository {
	return &ShortUrl{
		*postgres,
	}
}

func (i *ShortUrl) Create(url string) (uint64, error) {
	id, err := i.FindBy(url)
	if err == nil {
		return id, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	res := &entity.ShortUrl{Url: url}
	if err := i.DB.Create(res).Error; err != nil {
		return 0, err
	}
	return res.ID, nil
}

func (i *ShortUrl) FindBy(url string) (uint64, error) {
	res := &entity.ShortUrl{}
	if err := i.DB.Where("url = ?", url).First(res).Error; err != nil {
		return 0, err
	}
	return res.ID, nil
}

func (i *ShortUrl) Find(id uint64) (string, error) {
	res := &entity.ShortUrl{}
	if err := i.DB.First(&res, id).Error; err != nil {
		return "", err
	}
	return res.Url, nil
}

func (i *ShortUrl) Delete(id uint64) (uint64, error) {
	res := i.DB.Where("id = ?", id).Delete(&entity.ShortUrl{})
	return uint64(res.RowsAffected), res.Error
}
