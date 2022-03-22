package services

import (
	dto "go-rest/internal/application/dto/users"
	"go-rest/internal/domain/repositories"
)

type UsersService struct {
	repository repositories.IUsersRepository
}

func NewUsersService(r repositories.IUsersRepository) IUsersService {
	return &UsersService{r}
}

func (us *UsersService) GetById(id int) *dto.UserDto {
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
