package api

import (
	"github.com/gin-gonic/gin"
	"go-rest/internal/api/http/controllers"
	"go-rest/internal/api/http/handlers"
	"go-rest/internal/api/http/routers"
	"go-rest/internal/application/contracts"
	"go-rest/internal/config"
	"go-rest/internal/shared/models/di"
	"go.uber.org/dig"
	"net/http"
)

type AppRouters struct {
	dig.In

	Routers []contracts.IRouteBinder `group:"routers"`
}

func BindRouter(configuration *config.Configuration) []di.ServiceDescriptor {
	r := NewAppRouter()
	r.AddGlobalRoutePrefix(configuration.RoutePrefix)

	bindRouter := func() http.Handler {
		return r.Router
	}

	bindRouterGroup := func() *gin.RouterGroup {
		return r.RouterGroup
	}

	return []di.ServiceDescriptor{
		{Service: bindRouter},
		{Service: bindRouterGroup},
	}
}

func BindHandlers() []di.ServiceDescriptor {
	return []di.ServiceDescriptor{
		{Service: handlers.NewUsersHandler},
	}

}

func BindControllers() []di.ServiceDescriptor {
	return []di.ServiceDescriptor{
		{Service: controllers.NewUsersController},
	}
}

func BindRouterGroups() []di.ServiceDescriptor {
	serviceDescriptors := []di.ServiceDescriptor{
		{Service: routers.NewUsersRouter},
	}

	for i := 0; i < len(serviceDescriptors); i++ {
		if serviceDescriptors[i].Options == nil {
			serviceDescriptors[i].Options = make([]dig.ProvideOption, 0)
		}
		serviceDescriptors[i].Options = append(serviceDescriptors[i].Options.([]dig.ProvideOption), dig.Group("routers"))
	}

	return serviceDescriptors
}

func AfterBuild(container *dig.Container) {
	container.Invoke(func(ar AppRouters) {
		for _, r := range ar.Routers {
			r.Bind()
		}
	})
}
