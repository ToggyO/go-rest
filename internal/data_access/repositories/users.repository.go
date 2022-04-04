package repositories

import (
	"go-rest/internal/application/contracts"
	db_context "go-rest/internal/data_access/db-context"
	entities "go-rest/internal/data_access/entities/users"
	"go-rest/internal/domain/models/users"
	"go-rest/internal/domain/repositories"
	"gorm.io/gorm"
)

type usersRepository struct {
	context *gorm.DB
	mapper  contracts.IMapper
}

func NewUsersRepository(c *db_context.AppDbContext, m contracts.IMapper) repositories.IUsersRepository {
	return &usersRepository{c.GetConnection(), m}
}

func (ur *usersRepository) GetById(id int) *users.UserModel {
	model := new(users.UserModel)

	var entity *entities.UserEntity
	ur.context.First(entity, id)

	if entity == nil {
		return nil
	}

	ur.mapper.Map(entity, model)
	return model
}

func (ur *usersRepository) Create(model *users.UserModel) *users.UserModel {
	var entity entities.UserEntity

	ur.mapper.Map(model, &entity)
	ur.context.Create(&entity)

	model.Id = entity.Id
	return model
}

func (ur *usersRepository) Update(model *users.UserModel) *users.UserModel {
	entity := new(entities.UserEntity)
	ur.context.First(entity, model.Id)

	ur.mapper.Map(model, entity)
	ur.context.Save(entity)

	return model
}

func (ur *usersRepository) Delete(id int) {
	entity := new(entities.UserEntity)
	ur.context.First(entity, id)

	if entity != nil {
		ur.context.Delete(entity)
	}
}
