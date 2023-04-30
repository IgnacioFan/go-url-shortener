package migration

import (
	"go-url-shortener/internal/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var v20230312 = &gormigrate.Migration{
	ID: "20230312",
	Migrate: func(tx *gorm.DB) error {
		// when table already exists, it just adds fields as columns
		type ShortUrl struct {
			gorm.Model
			Url string `gorm:"type:varchar(1024) not null"`
		}
		return tx.AutoMigrate(&ShortUrl{})
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable(&entity.ShortUrl{}); err != nil {
			return err
		}
		return nil
	},
}
