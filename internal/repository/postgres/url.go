package postgres

import (
	"errors"
	"go-url-shortener/internal/entity"
	"time"

	"gorm.io/gorm"
)

type UrlRepository interface {
	Create(url string) (uint64, error)
	FindBy(url string) (uint64, error)
}

type Url struct {
	DB *gorm.DB
}

func (u *Url) Create(url string) (uint64, error) {
	id, err := u.FindBy(url)
	if err == nil {
		return id, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	res := &entity.Url{Url: url, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := u.DB.Create(res).Error; err != nil {
		return 0, err
	}
	return res.ID, nil
}

func (u *Url) FindBy(url string) (uint64, error) {
	res := &entity.Url{}
	if err := u.DB.Where("url = ?", url).First(res).Error; err != nil {
		return 0, err
	}
	return res.ID, nil
}
