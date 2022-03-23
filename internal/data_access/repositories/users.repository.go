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

type usersRepository struct {
	connection *sql.DB
}

// TODO: uncomment
//func NewUsersRepository(c *sql.DB) repositories.IUsersRepository {
//	return &UsersRepository{c}
//}

func NewUsersRepository() repositories.IUsersRepository {
	return &usersRepository{}
}

func (ur *usersRepository) GetById(id int) *users.User {
	return user
}

func (ur *usersRepository) Create(entity *users.User) *users.User {
	return user
}

func (ur *usersRepository) Update(entity *users.User) *users.User {
	return user
}

func (ur *usersRepository) Delete(id int) *users.User {
	return user
}
