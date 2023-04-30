package migration

import "github.com/go-gormigrate/gormigrate/v2"

var Migrations = []*gormigrate.Migration{
	v20230312,
	v20230413,
}
