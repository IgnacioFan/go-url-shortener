package entity

import "time"

type ShortUrl struct {
	ID         uint64
	Url        string `gorm:"type:varchar(1024) not null"`
	Expired_At time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (s *ShortUrl) TableName() string {
	return "short_urls"
}

type ShortUrlRepository interface {
	Create(url string) (uint64, error)
	FindBy(url string) (uint64, error)
	Find(id uint64) (string, error)
	Delete(id uint64) error
}

type ShortUrlUsecase interface {
	Create(url string) (string, error)
	Redirect(url string) (string, error)
	Delete(code string) error
}
