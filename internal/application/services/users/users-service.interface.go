package services

import (
	dto "go-rest/internal/application/dto/users"
)

type IUsersService interface {
	GetById(id int) *dto.UserDto
	Create(obj *dto.CreateUserDto) *dto.UserDto
	Update(obj *dto.UpdateUserDto) *dto.UserDto
	Delete(id int)
}
