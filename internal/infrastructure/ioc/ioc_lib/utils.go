package ioc_lib

import (
	"go-rest/internal/shared/models/di"
	"go.uber.org/dig"
)

func HandleServiceDescriptors(container *dig.Container, serviceDescriptors []di.ServiceDescriptor) error {
	for _, sd := range serviceDescriptors {
		if sd.Options == nil {
			return container.Provide(sd.Service)
		}

		return container.Provide(sd.Service, sd.Options.([]dig.ProvideOption)...)
	}

	return nil
}
