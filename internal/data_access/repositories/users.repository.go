package repositories

import (
	"database/sql"
	"go-rest/internal/domain/models/users"
	"go-rest/internal/domain/repositories"
)

// TODO: delete
var user = &users.User{
	Id:       1,
	Name:     "Slava",
	Email:    "Ukrainin",
	Password: "123456",
}

type UsersRepository struct {
	connection *sql.DB
}

// TODO: uncomment
//func NewUsersRepository(c *sql.DB) repositories.IUsersRepository {
//	return &UsersRepository{c}
//}

func NewUsersRepository() repositories.IUsersRepository {
	return &UsersRepository{}
}

func (ur *UsersRepository) GetById(id int) *users.User {
	return user
}

func (ur *UsersRepository) Create(entity *users.User) *users.User {
	return user
}

func (ur *UsersRepository) Update(entity *users.User) *users.User {
	return user
}

func (ur *UsersRepository) Delete(id int) *users.User {
	return user
}
