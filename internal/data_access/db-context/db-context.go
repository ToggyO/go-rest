package db_context

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: check
type IDbConnectionFactory[TConnection interface{}] interface {
	Create() (TConnection, error)
}

// TODO: check
type IDbContext interface {
	GetDbConnection()
}

type GormDbContext struct {
	connection *gorm.DB
}

// TODO: inspect gorm.Config struct
func NewDbContext(connectionString string) (*GormDbContext, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	return &GormDbContext{db}, err
}
