package ioc_lib

import (
	api "go-rest/internal/api/http"
	"go-rest/internal/application"
	"go-rest/internal/data_access"
	"go-rest/internal/infrastructure/config"
	"go-rest/internal/infrastructure/services"

	"go.uber.org/dig"
)

func BuildDigIoc(configuration *config.Configuration) (*dig.Container, error) {
	container := dig.New()

	errors := []error{
		container.Provide(func() *config.Configuration { return configuration }),
		api.BindRouter(container),
		api.BindHandlers(container),
		api.BindControllers(container),
		api.BindRouterGroups(container),
		data_access.BindDataAccess(container),
		application.BindApplicationServices(container),
		services.BindInfrastructure(container),
	}

	for _, val := range errors {
		if val != nil {
			return nil, val
		}
	}

	collection := []func(c *dig.Container){
		api.AfterBuild,
	}

	for _, f := range collection {
		f(container)
	}

	return container, nil
}
