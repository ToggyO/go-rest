package application

import (
	services "go-rest/internal/application/services/users"
	"go-rest/internal/infrastructure/ioc/ioc_utils"

	"go.uber.org/dig"
)

func BindApplicationServices(c *dig.Container) error {
	serviceDescriptors := []ioc_utils.ServiceDescriptor{
		{Service: services.NewUsersService},
	}

	return ioc_utils.HandleServiceDescriptors(c, serviceDescriptors)
}
