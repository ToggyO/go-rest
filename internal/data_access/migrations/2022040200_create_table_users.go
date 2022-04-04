package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	entities "go-rest/internal/data_access/entities/users"
	"gorm.io/gorm"
)

var CreateUsersTableMigration = gormigrate.Migration{
	ID: "2022040200",
	Migrate: func(db *gorm.DB) error {
		return db.Migrator().CreateTable(&entities.UserEntity{})
	},
	Rollback: func(db *gorm.DB) error {
		return db.Migrator().DropTable(&entities.UserEntity{})
	},
}
