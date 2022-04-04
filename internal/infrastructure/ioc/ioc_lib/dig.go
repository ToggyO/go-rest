package ioc_lib

import (
	api "go-rest/internal/api/http"
	"go-rest/internal/application"
	"go-rest/internal/config"
	"go-rest/internal/data_access"
	"go-rest/internal/infrastructure/services"
	"go-rest/internal/shared/models/di"

	"go.uber.org/dig"
)

func BuildDigIoc(configuration *config.Configuration) (*dig.Container, error) {
	container := dig.New()
	container.Provide(func() *config.Configuration { return configuration })

	serviceDescriptors := [][]di.ServiceDescriptor{
		api.BindRouter(configuration),
		api.BindHandlers(),
		api.BindControllers(),
		api.BindRouterGroups(),
		data_access.BindDataAccess(),
		application.BindApplicationServices(),
		services.BindInfrastructure(),
	}

	for _, sd := range serviceDescriptors {
		if err := HandleServiceDescriptors(container, sd); err != nil {
			return nil, err
		}
	}

	collection := []func(c *dig.Container){
		api.AfterBuild,
		data_access.AfterBuild,
	}

	for _, f := range collection {
		f(container)
	}

	return container, nil
}
