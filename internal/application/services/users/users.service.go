package services

import (
	dto "go-rest/internal/application/dto/users"
	"go-rest/internal/domain/models/users"
	"go-rest/internal/domain/repositories"
)

type usersService struct {
	repository repositories.IUsersRepository
}

func NewUsersService(r repositories.IUsersRepository) IUsersService {
	return &usersService{r}
}

func (us *usersService) GetById(id int) *dto.UserDto {
	model := us.repository.GetById(id)
	if model == nil {
		return nil
	}
	// TODO: add automapper
	return &dto.UserDto{
		Id:    model.Id,
		Name:  model.Name,
		Email: model.Email,
	}
}

func (us *usersService) Create(obj *dto.CreateUserDto) *dto.UserDto {
	// TODO: add automapper
	// TODO: hash password
	model := &users.User{
		Name:  obj.Name,
		Email: obj.Email,
	}

	model = us.repository.Create(model)
	return &dto.UserDto{
		Id:    model.Id,
		Name:  model.Name,
		Email: model.Email,
	}
}
