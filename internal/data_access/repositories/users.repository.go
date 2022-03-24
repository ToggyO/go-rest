package repositories

import (
	"database/sql"
	"go-rest/internal/domain/models/users"
	"go-rest/internal/domain/repositories"
)

// TODO: delete
var user = &users.User{
	Id:    1,
	Name:  "Slava",
	Email: "Ukrainin",
}

type usersRepository struct {
	connection *sql.DB
	users      []*users.User
}

// TODO: uncomment
//func NewUsersRepository(c *sql.DB) repositories.IUsersRepository {
//	return &UsersRepository{c}
//}

func NewUsersRepository() repositories.IUsersRepository {
	return &usersRepository{
		users: []*users.User{
			{
				Id:    1,
				Name:  "Slava",
				Email: "Ukrainin",
			},
			{
				Id:    2,
				Name:  "Vasia",
				Email: "Pimshin",
			},
			{
				Id:    3,
				Name:  "Sasha",
				Email: "Sosagun",
			},
		},
	}
}

func (ur *usersRepository) GetById(id int) *users.User {
	var user *users.User
	for _, v := range ur.users {
		if v.Id == id {
			user = v
		}
	}
	return user
}

func (ur *usersRepository) Create(model *users.User) *users.User {
	_, max := ur.findMaxAndMinId()
	model.Id = max + 1
	ur.users = append(ur.users, model)
	return model
}

func (ur *usersRepository) Update(model *users.User) *users.User {
	return user
}

func (ur *usersRepository) Delete(id int) *users.User {
	return user
}

func (ur *usersRepository) findMaxAndMinId() (min int, max int) {
	if len(ur.users) == 0 {
		min = 0
		max = 0
		return min, max
	}

	min = ur.users[0].Id
	max = ur.users[0].Id

	for _, v := range ur.users {
		if v.Id > max {
			max = v.Id
		}
		if v.Id < min {
			min = v.Id
		}
	}

	return min, max
}
