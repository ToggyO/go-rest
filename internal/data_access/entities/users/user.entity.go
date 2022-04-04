package entities

import "time"

type UserEntity struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Hash      string
	Salt      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u UserEntity) TableName() string {
	return "users"
}
