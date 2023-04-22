package postgres

import (
	"go-url-shortener/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConn() *gorm.DB {
	config, err := config.New()
	db, err := gorm.Open(postgres.Open(config.Postgres.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
