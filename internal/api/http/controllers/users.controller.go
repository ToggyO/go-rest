package controllers

import (
	"go-rest/internal/api/http/handlers"
	dto "go-rest/internal/application/dto/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IUsersController interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type usersController struct {
	handler handlers.IUsersHandler
}

func (uc *usersController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// TODO: return error response
	}

	result := uc.handler.GetById(id)
	ctx.JSON(result.HttpStatusCode, result)
}

func (uc *usersController) Create(ctx *gin.Context) {
	body := &dto.CreateUserDto{}
	err := ctx.Bind(body)
	if err != nil {
		// TODO: return error response
	}

	result := uc.handler.Create(body)
	ctx.JSON(result.HttpStatusCode, result)
}

func NewUsersController(h handlers.IUsersHandler) IUsersController {
	return &usersController{h}
}
