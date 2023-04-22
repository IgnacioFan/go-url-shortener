package postgres

import (
	"go-url-shortener/deployment/migration"

	"github.com/go-gormigrate/gormigrate/v2"
)

func NewMigrate() error {
	if err := gormigrate.New(InitConn(), gormigrate.DefaultOptions, migration.Migrations).Migrate(); err != nil {
		return err
	}
	return nil
}
