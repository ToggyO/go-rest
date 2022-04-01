package ioc_lib

import (
	"go-rest/internal/shared/models/di"
	"go.uber.org/dig"
)

func HandleServiceDescriptors(container *dig.Container, serviceDescriptors []di.ServiceDescriptor) error {
	var err error
	for _, sd := range serviceDescriptors {
		if sd.Options == nil {
			err = container.Provide(sd.Service)
			continue
		}

		err = container.Provide(sd.Service, sd.Options.([]dig.ProvideOption)...)
	}

	return err
}
