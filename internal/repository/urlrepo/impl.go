package urlrepo

import (
	"errors"
	"go-url-shortener/internal/entity"
	"go-url-shortener/internal/repository"
	"time"

	"gorm.io/gorm"
)

type Impl struct {
	DB *gorm.DB
}

func NewUrlRepository(db *gorm.DB) repository.UrlRepository {
	return &Impl{
		DB: db,
	}
}

func (i *Impl) Create(url string) (uint64, error) {
	id, err := i.FindBy(url)
	if err == nil {
		return id, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	res := &entity.Url{Url: url, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := i.DB.Create(res).Error; err != nil {
		return 0, err
	}
	return res.ID, nil
}

func (i *Impl) FindBy(url string) (uint64, error) {
	res := &entity.Url{}
	if err := i.DB.Where("url = ?", url).First(res).Error; err != nil {
		return 0, err
	}
	return res.ID, nil
}

func (i *Impl) Find(id uint64) (string, error) {
	res := &entity.Url{}
	if err := i.DB.First(&res, id).Error; err != nil {
		return "", err
	}
	return res.Url, nil
}
