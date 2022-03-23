package services

import (
	dto "go-rest/internal/application/dto/users"
	"go-rest/internal/domain/repositories"
)

type usersService struct {
	repository repositories.IUsersRepository
}

func NewUsersService(r repositories.IUsersRepository) IUsersService {
	return &usersService{r}
}

func (us *usersService) GetById(id int) *dto.UserDto {
	entity := us.repository.GetById(id)
	if entity == nil {
		// TODO: handle
	}
	// TODO: add automapper
	return &dto.UserDto{
		Id:    entity.Id,
		Name:  entity.Name,
		Email: entity.Email,
	}
}
