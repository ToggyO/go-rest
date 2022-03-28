package data_access

import (
	"go-rest/internal/data_access/repositories"
	"go-rest/internal/infrastructure/ioc/ioc_lib"
	"go.uber.org/dig"
)

func BindDataAccess(container *dig.Container) error {
	serviceDescriptors := []ioc_lib.ServiceDescriptor{
		{Service: repositories.NewUsersRepository},
	}

	return ioc_lib.HandleServiceDescriptors(container, serviceDescriptors)
}
