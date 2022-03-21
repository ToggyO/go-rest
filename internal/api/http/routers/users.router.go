package routers

import (
	"github.com/gin-gonic/gin"
	"go-rest/internal/api/http/controllers"
)

type UsersRouter struct {
	usersController controllers.UsersController
	handler         *gin.Engine
}

func (r *UsersRouter) Bind() {
	r.handler.Group("/users")
	{
		r.handler.GET(":id", r.usersController.GetById)
	}
}

func NewUsersRouter(uc controllers.UsersController, h *gin.Engine) UsersRouter {
	return UsersRouter{
		usersController: uc,
		handler:         h,
	}
}
