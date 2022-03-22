package data_access

import (
	"go-rest/internal/data_access/repositories"
	"go-rest/internal/infrastructure/ioc/ioc_utils"
	"go.uber.org/dig"
)

func BindDataAccess(container *dig.Container) error {
	serviceDescriptors := []ioc_utils.ServiceDescriptor{
		{Service: repositories.NewUsersRepository},
	}

	return ioc_utils.HandleServiceDescriptors(container, serviceDescriptors)
}
