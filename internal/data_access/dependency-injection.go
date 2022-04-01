package data_access

import (
	db_context "go-rest/internal/data_access/db-context"
	"go-rest/internal/data_access/repositories"
	"go-rest/internal/shared/models/di"
)

func BindDataAccess() []di.ServiceDescriptor {
	return []di.ServiceDescriptor{
		{Service: db_context.NewDbContext},
		{Service: repositories.NewUsersRepository},
	}
}
