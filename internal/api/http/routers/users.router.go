package routers

import (
	"github.com/gin-gonic/gin"
	"go-rest/internal/api/http/controllers"
	"go-rest/internal/application/contracts"
)

type usersRouter struct {
	usersController controllers.IUsersController
	handler         *gin.RouterGroup
}

func NewUsersRouter(uc controllers.IUsersController, h *gin.RouterGroup) contracts.IRouteBinder {
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
	}
}
