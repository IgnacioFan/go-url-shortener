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
