package migration

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var v20230413 = &gormigrate.Migration{
	ID: "20230413",
	Migrate: func(tx *gorm.DB) error {
		// when table already exists, it just adds fields as columns
		type ShortUrl struct {
			Expired_At time.Time
		}
		return tx.AutoMigrate(&ShortUrl{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropColumn("short_urls", "expired_at")
	},
}
