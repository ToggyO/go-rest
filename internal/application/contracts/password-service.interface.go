package contracts

import "go-rest/internal/domain/models/users"

type IPasswordService interface {
	// CreatePassword - creates password
	CreatePassword(password string, salt string) *users.PasswordModel

	// CreatePasswordAndGenerateSalt - creates password with generated salt
	CreatePasswordAndGenerateSalt(password string) *users.PasswordModel

	// VerifyPassword - verify password
	VerifyPassword(currentPassword *users.PasswordModel, password string) bool

	// GenerateSalt - generates random salt
	GenerateSalt() string
}
