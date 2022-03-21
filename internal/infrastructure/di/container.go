package ioc

import (
	api "go-rest/internal/api/http"
	"go-rest/internal/data_access"
	"go.uber.org/dig"
)

func BuildIoc() (*dig.Container, error) {
	container := dig.New()

	errors := []error{
		container.Provide(api.BindRouter),
		api.BindControllers(container),
		api.BindRouterGroups(container),
		data_access.BindDataAccess(container),
	}

	for _, val := range errors {
		if val != nil {
			return nil, val
		}
	}

	return container, nil
}
