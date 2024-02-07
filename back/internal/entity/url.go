package entity

import (
	"time"
)

type Url struct {
	ID        uint      `gorm:"primaryKey"`
	LongURL	  string		`gorm:"not null"`
	ShortURL  string		`gorm:"type:char(6) not null"`
  CreatedAt time.Time
  UpdatedAt time.Time
}
