package migration

import (
	"go-url-shortener/internal/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var v20220312 = &gormigrate.Migration{
	ID: "20220312",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&entity.Url{}); err != nil {
			return err
		}
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable(&entity.Url{}); err != nil {
			return err
		}
		return nil
	},
}
