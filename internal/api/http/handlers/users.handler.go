package handlers

import (
	dto "go-rest/internal/application/dto/users"
	services "go-rest/internal/application/services/users"
	"go-rest/internal/shared/constants/errors/error_codes/global"
	"go-rest/internal/shared/models/response"
)

type IUsersHandler interface {
	GetById(id int) *response.Response[*dto.UserDto]
}

type usersHandler struct {
	service services.IUsersService
}

func NewUsersHandler(s services.IUsersService) IUsersHandler {
	return &usersHandler{s}
}

func (uh *usersHandler) GetById(id int) *response.Response[*dto.UserDto] {
	r := response.NewDefaultResponse[*dto.UserDto]()

	user := uh.service.GetById(id)
	if user == nil {
		// TODO: into constants
		return r.ToErrorResponse(404, global.NotFound, "Not found", r.DefaultValidationError())
	}

	r.SuccessResponse.Data = user
	return r

}
