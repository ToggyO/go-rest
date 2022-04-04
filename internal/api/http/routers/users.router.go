package routers

import (
	"github.com/gin-gonic/gin"
	"go-rest/internal/api/http/controllers"
	"go-rest/internal/application/contracts"
)

type usersRouter struct {
	usersController *controllers.UsersController
	handler         *gin.RouterGroup
}

func NewUsersRouter(uc *controllers.UsersController, h *gin.RouterGroup) contracts.IRouteBinder {
	return &usersRouter{
		usersController: uc,
		handler:         h,
	}
}

func (r *usersRouter) Bind() {
	users := r.handler.Group("/users")
	{
		users.GET("/:id", r.usersController.GetById)
		users.POST("", r.usersController.Create)
		users.PATCH("", r.usersController.Update)
		users.DELETE("/:id", r.usersController.Delete)
	}
}
