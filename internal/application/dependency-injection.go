package application

import (
	services "go-rest/internal/application/services/users"
	"go-rest/internal/shared/models/di"
)

func BindApplicationServices() []di.ServiceDescriptor {
	return []di.ServiceDescriptor{
		{Service: services.NewUsersService},
	}
}
