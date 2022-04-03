package repositories

import (
	db_context "go-rest/internal/data_access/db-context"
	entities "go-rest/internal/data_access/entities/users"
	"go-rest/internal/domain/models/users"
	"go-rest/internal/domain/repositories"
	"gorm.io/gorm"
)

// TODO: delete
var user = &users.UserModel{
	Id:    1,
	Name:  "Slava",
	Email: "Ukrainin",
}

type usersRepository struct {
	context *gorm.DB
	//users      []*users.UserModel
}

func NewUsersRepository(c *db_context.AppDbContext) repositories.IUsersRepository {
	return &usersRepository{c.GetConnection()}
}

//func NewUsersRepository() repositories.IUsersRepository {
//	return &usersRepository{
//		users: []*users.UserModel{
//			{
//				Id:    1,
//				Name:  "Slava",
//				Email: "Ukrainin",
//			},
//			{
//				Id:    2,
//				Name:  "Vasia",
//				Email: "Pimshin",
//			},
//			{
//				Id:    3,
//				Name:  "Sasha",
//				Email: "Sosagun",
//			},
//		},
//	}
//}

func (ur *usersRepository) GetById(id int) *users.UserModel {
	// TODO: add automapper
	var user *users.UserModel
	entity := entities.UserEntity{}
	ur.context.First(user, id)
	return user
}

func (ur *usersRepository) Create(model *users.UserModel) *users.UserModel {

	return model
}

func (ur *usersRepository) Update(model *users.UserModel) *users.UserModel {
	return user
}

func (ur *usersRepository) Delete(id int) *users.UserModel {
	return user
}

//func (ur *usersRepository) findMaxAndMinId() (min int, max int) {
//	if len(ur.users) == 0 {
//		min = 0
//		max = 0
//		return min, max
//	}
//
//	min = ur.users[0].Id
//	max = ur.users[0].Id
//
//	for _, v := range ur.users {
//		if v.Id > max {
//			max = v.Id
//		}
//		if v.Id < min {
//			min = v.Id
//		}
//	}
//
//	return min, max
//}
