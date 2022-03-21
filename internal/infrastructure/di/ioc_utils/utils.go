package ioc_utils

import (
	"go.uber.org/dig"
)

type ServiceDescriptor struct {
	Service interface{}
	Options []dig.ProvideOption
}

func HandleServiceDescriptors(container *dig.Container, serviceDescriptors []ServiceDescriptor) error {
	for _, sd := range serviceDescriptors {
		if err := container.Provide(sd.Service, sd.Options...); err != nil {
			return err
		}
	}

	return nil
}
