package migration

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var Migrations = []*gormigrate.Migration{
  {
    ID: "20231230",
    Migrate: func(tx *gorm.DB) error {
      type Url struct {
        ID        uint      `gorm:"primaryKey"`
        LongURL	  string		`gorm:"not null"`
        ShortURL  string		`gorm:"type:char(6) not null"`
        CreatedAt time.Time
        UpdatedAt time.Time
      }
      return tx.AutoMigrate(&Url{})
    },
    Rollback: func(tx *gorm.DB) error {
      return tx.Migrator().DropTable("urls")
    },
  },
}
