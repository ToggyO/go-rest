package repositories

import "go-rest/internal/domain/models/users"

type IUsersRepository interface {
	IBaseRepository[users.User]
}
