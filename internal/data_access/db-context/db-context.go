package db_context

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppDbContext struct {
	connection *gorm.DB
}

// TODO: inspect gorm.Config struct
func NewDbContext(connectionString string) (*AppDbContext, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	return &AppDbContext{db}, err
}

func (context *AppDbContext) GetConnection() *gorm.DB {
	return context.connection
}
