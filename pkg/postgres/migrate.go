package postgres

import (
	"go-url-shortener/deployment/migration"

	"github.com/go-gormigrate/gormigrate/v2"
)

func (p *Postgres) NewMigrate() error {
	if err := gormigrate.New(p.DB, gormigrate.DefaultOptions, migration.Migrations).Migrate(); err != nil {
		return err
	}
	return nil
}
