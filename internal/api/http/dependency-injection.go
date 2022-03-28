package api

import (
	"go-rest/internal/api/http/controllers"
	"go-rest/internal/api/http/handlers"
	"go-rest/internal/api/http/routers"
	"go-rest/internal/application/contracts"
	"go-rest/internal/infrastructure/ioc/ioc_lib"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type AppRouters struct {
	dig.In

	Routers []contracts.IRouteBinder `group:"routers"`
}

func BindRouter(container *dig.Container) error {
	r := NewAppRouter()
	r.AddGlobalRoutePrefix("/api")

	bindRouter := func() http.Handler {
		return r.Router
	}

	bindRouterGroup := func() *gin.RouterGroup {
		return r.RouterGroup
	}

	serviceDescriptors := []ioc_lib.ServiceDescriptor{
		{Service: bindRouter},
		{Service: bindRouterGroup},
	}

	return ioc_lib.HandleServiceDescriptors(container, serviceDescriptors)
}

func BindHandlers(container *dig.Container) error {
	serviceDescriptors := []ioc_lib.ServiceDescriptor{
		{Service: handlers.NewUsersHandler},
	}

	return ioc_lib.HandleServiceDescriptors(container, serviceDescriptors)
}

func BindControllers(container *dig.Container) error {
	serviceDescriptors := []ioc_lib.ServiceDescriptor{
		{Service: controllers.NewUsersController},
	}

	return ioc_lib.HandleServiceDescriptors(container, serviceDescriptors)
}

func BindRouterGroups(container *dig.Container) error {
	serviceDescriptors := []ioc_lib.ServiceDescriptor{
		{Service: routers.NewUsersRouter},
	}

	for i := 0; i < len(serviceDescriptors); i++ {
		serviceDescriptors[i].Options = append(serviceDescriptors[i].Options, dig.Group("routers"))
	}

	return ioc_lib.HandleServiceDescriptors(container, serviceDescriptors)
}

func AfterBuild(container *dig.Container) {
	container.Invoke(func(ar AppRouters) {
		for _, r := range ar.Routers {
			r.Bind()
		}
	})
}
