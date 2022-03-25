package services

import (
	"go-rest/internal/infrastructure/ioc/ioc_utils"
	"go.uber.org/dig"
)

func BindInfrastructure(container *dig.Container) error {
	serviceDescriptor := []ioc_utils.ServiceDescriptor{
		{Service: NewPasswordService},
	}

	return ioc_utils.HandleServiceDescriptors(container, serviceDescriptor)
}
