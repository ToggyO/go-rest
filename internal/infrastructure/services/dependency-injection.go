package services

import (
	"go-rest/internal/infrastructure/services/logger"
	"go-rest/internal/infrastructure/services/password"
	"go-rest/internal/shared/models/di"
)

func BindInfrastructure() []di.ServiceDescriptor {
	return []di.ServiceDescriptor{
		{Service: logger.NewLoggerService},
		{Service: password.NewPasswordService},
	}
}
