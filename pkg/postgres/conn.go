package postgres

import (
	"go-url-shortener/deployment/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres() (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(env.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Postgres{DB: db}, nil
}
