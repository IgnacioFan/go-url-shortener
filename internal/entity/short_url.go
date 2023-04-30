package entity

import (
	"time"

	"gorm.io/gorm"
)

type ShortUrl struct {
	ID        uint64
	Url       string `gorm:"type:varchar(1024) not null"`
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (s *ShortUrl) TableName() string {
	return "short_urls"
}

type ShortUrlRepository interface {
	Create(url string) (uint64, error)
	FindBy(url string) (uint64, error)
	Find(id uint64) (string, error)
	Delete(id uint64) (uint64, error)
}

type ShortUrlUsecase interface {
	Create(url string) (string, error)
	Redirect(url string) (string, error)
	Delete(code string) error
}
