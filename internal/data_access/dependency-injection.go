package data_access

import (
	"go-rest/internal/data_access/repositories"
	"go-rest/internal/shared/models/di"
)

func BindDataAccess() []di.ServiceDescriptor {
	return []di.ServiceDescriptor{
		{Service: repositories.NewUsersRepository},
	}
}
