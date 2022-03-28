package password

import (
	"crypto/sha512"
	"encoding/base64"
	"go-rest/internal/application/contracts"
	"go-rest/internal/domain/models/users"
	"math/rand"
)

type passwordService struct{}

func NewPasswordService() contracts.IPasswordService {
	return &passwordService{}
}

func (ps *passwordService) CreatePassword(password string, salt string) *users.PasswordModel {
	passwordBytes := []byte(password)
	passwordBytes = append(passwordBytes, []byte(salt)...)

	sha512Hasher := sha512.New()
	sha512Hasher.Write(passwordBytes)

	return &users.PasswordModel{
		Salt: salt,
		Hash: base64.URLEncoding.EncodeToString(sha512Hasher.Sum(nil)),
	}
}

func (ps *passwordService) CreatePasswordAndGenerateSalt(password string) *users.PasswordModel {
	return ps.CreatePassword(password, ps.GenerateSalt())
}

func (ps *passwordService) VerifyPassword(currentPassword *users.PasswordModel, password string) bool {
	return ps.CreatePassword(password, currentPassword.Salt).Hash == currentPassword.Hash
}

func (ps *passwordService) GenerateSalt() string {
	return base64.StdEncoding.EncodeToString(ps.generateRandomBytes())
}

func (ps *passwordService) generateRandomBytes() []byte {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return bytes
}
