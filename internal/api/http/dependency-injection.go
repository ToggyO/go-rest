package api

import (
	"go-rest/internal/api/http/controllers"
	"go-rest/internal/api/http/handlers"
	"go-rest/internal/api/http/routers"
	"go-rest/internal/application/contracts"
	"go-rest/internal/infrastructure/ioc/ioc_utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type AppRouters struct {
	dig.In

	Routers []contracts.IRouteBinder `group:"routers"`
}

func BindRouter() http.Handler {
	return gin.Default()
}

func BindHandlers(container *dig.Container) error {
	serviceDescriptors := []ioc_utils.ServiceDescriptor{
		{Service: handlers.NewUsersHandler},
	}

	return ioc_utils.HandleServiceDescriptors(container, serviceDescriptors)
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
