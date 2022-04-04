package controllers

import (
	"go-rest/internal/api/http/handlers"
	dto "go-rest/internal/application/dto/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	handler handlers.IUsersHandler
}

func NewUsersController(h handlers.IUsersHandler) *UsersController {
	return &UsersController{h}
}

func (uc *UsersController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// TODO: return error response
	}

	result := uc.handler.GetById(id)
	ctx.JSON(result.HttpStatusCode, result)
}

func (uc *UsersController) Create(ctx *gin.Context) {
	body := &dto.CreateUserDto{}
	err := ctx.Bind(body)
	if err != nil {
		// TODO: return error response
	}

	result := uc.handler.Create(body)
	ctx.JSON(result.HttpStatusCode, result)
}

func (uc *UsersController) Update(ctx *gin.Context) {
	body := &dto.UpdateUserDto{}
	err := ctx.Bind(body)
	if err != nil {
		// TODO: return error response
	}

	result := uc.handler.Update(body)
	ctx.JSON(result.HttpStatusCode, result)
}

func (uc *UsersController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// TODO: return error response
	}

	result := uc.handler.Delete(id)
	ctx.JSON(result.HttpStatusCode, result)
}
