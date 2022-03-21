package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rest/internal/domain/models/users"
	"net/http"
	"strconv"
)

type UsersController struct{}

func (u *UsersController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// todo: return error response
	}

	user := users.User{
		Id:       id,
		Name:     "Slava",
		Email:    "Ukrainin",
		Password: "123456",
	}

	ctx.JSON(http.StatusOK, user)
}

func NewUsersController() *UsersController {
	return &UsersController{}
}
