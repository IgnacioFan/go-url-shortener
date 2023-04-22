package entity

import "time"

type Url struct {
	ID         uint64
	Url        string `gorm:"type:varchar(1024) not null"`
	Expired_At time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (s *Url) TableName() string {
	return "urls"
}

type ShortUrlRepository interface {
	Create(url string) (uint64, error)
	FindBy(url string) (uint64, error)
	Find(id uint64) (string, error)
}

type ShortUrlUsecase interface {
	Create(url string) (string, error)
	Redirect(url string) (string, error)
}
