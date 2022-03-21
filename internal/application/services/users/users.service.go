package services

import (
	dto "go-rest/internal/application/dto/users"
	"go-rest/internal/domain/repositories"
)

type IUsersService interface {
	GetById(id int) *dto.UserDto
}

type UsersService struct {
	repository repositories.IUsersRepository
}

func (us *UsersService) GetById(id int) *dto.UserDto {
	entity := us.repository.GetById(id)
	if entity == nil {
		// TODO: handle
	}
	return &dto.UserDto {
		Id: entity
	}
}

