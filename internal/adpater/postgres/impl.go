package postgres

import (
	"go-url-shortener/deployment/env"
	"go-url-shortener/internal/migration"

	"github.com/go-gormigrate/gormigrate/v2"
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

func (p *Postgres) NewMigrate() error {
	if err := gormigrate.New(p.DB, gormigrate.DefaultOptions, migration.Migrations).Migrate(); err != nil {
		return err
	}
	return nil
}
