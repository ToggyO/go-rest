package application

import (
	services "go-rest/internal/application/services/users"
	"go-rest/internal/infrastructure/ioc/ioc_lib"

	"go.uber.org/dig"
)

func BindApplicationServices(c *dig.Container) error {
	serviceDescriptors := []ioc_lib.ServiceDescriptor{
		{Service: services.NewUsersService},
	}

	return ioc_lib.HandleServiceDescriptors(c, serviceDescriptors)
}
