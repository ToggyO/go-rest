package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-rest/internal/api/http/controllers"
	"go-rest/internal/api/http/routers"
	"go.uber.org/dig"
	"net/http"
)

func BindRouter() http.Handler {
	return gin.Default()
}

func BindControllers() wire.ProviderSet {
	return wire.NewSet(controllers.NewUsersController)
}

func BindRouterGroups() wire.ProviderSet {
	return wire.NewSet(routers.NewUsersRouter)
}

func K() {
	container := dig.Container{}
	dig.As(BindRouter)
	container.Provide(BindRouter)
}
