package postgres

import (
	"go-url-shortener/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func InitPostgres(config *config.Config) *Postgres {
	db, err := gorm.Open(postgres.Open(config.Postgres.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Postgres{
		DB: db,
	}
}
