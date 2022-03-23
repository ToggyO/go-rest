package ioc

import (
	api "go-rest/internal/api/http"
	"go-rest/internal/application"
	"go-rest/internal/data_access"

	"go.uber.org/dig"
)

// TODO: turn into API layer mb or just into separated folder
func BuildIoc() (*dig.Container, error) {
	container := dig.New()

	errors := []error{
		container.Provide(api.BindRouter),
		api.BindHandlers(container),
		api.BindControllers(container),
		api.BindRouterGroups(container),
		data_access.BindDataAccess(container),
		application.BindApplicationServices(container),
	}

	for _, val := range errors {
		if val != nil {
			return nil, val
		}
	}

	return container, nil
}
