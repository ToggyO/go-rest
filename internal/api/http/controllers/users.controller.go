package controllers

import (
	"go-rest/internal/api/http/handlers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type usersController struct {
	handler handlers.IUsersHandler
}

func (uc *usersController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// TODO: return error response
	}

	result := uc.handler.GetById(id)
	ctx.JSON(http.StatusOK, result)
}

func NewUsersController(h handlers.IUsersHandler) *usersController {
	return &usersController{h}
}
