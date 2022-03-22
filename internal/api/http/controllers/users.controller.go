package controllers

import (
	services "go-rest/internal/application/services/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	service services.IUsersService
}

func (uc *UsersController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// todo: return error response
	}

	user := uc.service.GetById(id)

	ctx.JSON(http.StatusOK, user)
}

func NewUsersController(s services.IUsersService) *UsersController {
	return &UsersController{s}
}
