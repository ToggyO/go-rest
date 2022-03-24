package handlers

import (
	dto "go-rest/internal/application/dto/users"
	services "go-rest/internal/application/services/users"
	"go-rest/internal/shared/models/responses"
)

type IUsersHandler interface {
	GetById(id int) *responses.Response[*dto.UserDto]
	Create(obj *dto.CreateUserDto) *responses.Response[*dto.UserDto]
}

type usersHandler struct {
	service services.IUsersService
}

func NewUsersHandler(s services.IUsersService) IUsersHandler {
	return &usersHandler{s}
}

func (uh *usersHandler) GetById(id int) *responses.Response[*dto.UserDto] {
	user := uh.service.GetById(id)
	if user == nil {
		return responses.NewNotFoundErrorResponse[*dto.UserDto]()
	}

	r := responses.NewResponse[*dto.UserDto]()
	r.SuccessResponse.Data = user
	return r

}

func (uh *usersHandler) Create(obj *dto.CreateUserDto) *responses.Response[*dto.UserDto] {
	user := uh.service.Create(obj)

	r := responses.NewResponse[*dto.UserDto]()
	r.SuccessResponse.Data = user

	return r
}
