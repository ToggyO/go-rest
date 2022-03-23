package routers

import (
	"go-rest/internal/api/http/controllers"
	"go-rest/internal/application/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usersRouter struct {
	usersController *controllers.UsersController
	handler         *gin.Engine
}

func NewUsersRouter(uc *controllers.UsersController, h http.Handler) contracts.IRouteBinder {
	return &usersRouter{
		usersController: uc,
		handler:         h.(*gin.Engine),
	}
}

func (r *usersRouter) Bind() {
	users := r.handler.Group("/users")
	{
		users.GET("/:id", r.usersController.GetById)
	}
}
