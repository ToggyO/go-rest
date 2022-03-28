package ioc

import (
	"go-rest/internal/infrastructure/config"
	"go-rest/internal/infrastructure/ioc/ioc_lib"
	"go.uber.org/dig"
)

// TODO: can be user after generic types for struct methods will be allowed
var container = dig.New()

func NewGenericIoc(configuration *config.Configuration) (*dig.Container, error) {
	container, err := ioc_lib.BuildDigIoc(configuration)
	if err != nil {
		return nil, err
	}
	return container, nil
}

func GetService[TService any]() (TService, error) {
	var result TService
	err := container.Invoke(func(k TService) {
		result = k
	})
	return result, err
}

func SetService[TService any](sk TService) error {
	err := container.Provide(func() TService {
		return sk
	})
	return err
}
