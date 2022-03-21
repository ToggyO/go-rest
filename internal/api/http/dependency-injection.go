package api

import (
	"github.com/gin-gonic/gin"
	"go-rest/internal/api/http/controllers"
	"go-rest/internal/api/http/routers"
	"go-rest/internal/application/contracts"
	"go-rest/internal/infrastructure/di/ioc_utils"
	"go.uber.org/dig"
	"net/http"
)

type AppRouters struct {
	dig.In

	Routers []contracts.IRouteBinder `group:"routers"`
}

func BindRouter() http.Handler {
	return gin.Default()
}

func BindControllers(container *dig.Container) error {
	serviceDescriptors := []ioc_utils.ServiceDescriptor{
		{Service: controllers.NewUsersController},
	}

	return ioc_utils.HandleServiceDescriptors(container, serviceDescriptors)
}

func BindRouterGroups(container *dig.Container) error {
	serviceDescriptors := []ioc_utils.ServiceDescriptor{
		{Service: routers.NewUsersRouter},
	}

	for i := 0; i < len(serviceDescriptors); i++ {
		serviceDescriptors[i].Options = append(serviceDescriptors[i].Options, dig.Group("routers"))
	}

	return ioc_utils.HandleServiceDescriptors(container, serviceDescriptors)
}
