package model

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	ID        uint64
	ShortUrl  string `gorm:"not null"`
	LongUrl   string `gorm:"not null"`
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
