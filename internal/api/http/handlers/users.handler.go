package handlers

import (
	"go-rest/internal/api/http/validation"
	dto "go-rest/internal/application/dto/users"
	services "go-rest/internal/application/services/users"
	"go-rest/internal/shared/models/responses"
)

type IUsersHandler interface {
	GetById(id int) *responses.Response[*dto.UserDto]
	Create(obj *dto.CreateUserDto) *responses.Response[*dto.UserDto]
	Update(obj *dto.UpdateUserDto) *responses.Response[*dto.UserDto]
	Delete(id int) *responses.Response[interface{}]
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
	validationResponse := validation.ValidateModel[*dto.CreateUserDto, *dto.UserDto](obj)
	if !validationResponse.Success {
		return validationResponse
	}

	user := uh.service.Create(obj)

	r := responses.NewResponse[*dto.UserDto]()
	r.SuccessResponse.Data = user

	return r
}

func (uh *usersHandler) Update(obj *dto.UpdateUserDto) *responses.Response[*dto.UserDto] {
	validationResponse := validation.ValidateModel[*dto.UpdateUserDto, *dto.UserDto](obj)
	if !validationResponse.Success {
		return validationResponse
	}

	user := uh.service.Update(obj)

	r := responses.NewResponse[*dto.UserDto]()
	r.SuccessResponse.Data = user

	return r
}

func (uh *usersHandler) Delete(id int) *responses.Response[interface{}] {
	uh.service.Delete(id)
	return responses.NewResponse[interface{}]()
}
