package services

import (
	"go-rest/internal/infrastructure/ioc/ioc_lib"
	"go-rest/internal/infrastructure/services/logger"
	"go-rest/internal/infrastructure/services/password"
	"go.uber.org/dig"
)

func BindInfrastructure(container *dig.Container) error {
	serviceDescriptor := []ioc_lib.ServiceDescriptor{
		{Service: logger.NewLoggerService},
		{Service: password.NewPasswordService},
	}

	return ioc_lib.HandleServiceDescriptors(container, serviceDescriptor)
}
