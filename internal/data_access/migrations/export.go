package migrations

import "github.com/go-gormigrate/gormigrate/v2"

var MigrationCollection = []*gormigrate.Migration{
	&CreateUsersTableMigration,
}
