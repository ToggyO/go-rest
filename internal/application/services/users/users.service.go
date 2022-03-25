package services

import (
	"go-rest/internal/application/contracts"
	dto "go-rest/internal/application/dto/users"
	"go-rest/internal/domain/models/users"
	"go-rest/internal/domain/repositories"
)

type usersService struct {
	repository      repositories.IUsersRepository
	passwordService contracts.IPasswordService
}

func NewUsersService(r repositories.IUsersRepository, ps contracts.IPasswordService) IUsersService {
	return &usersService{r, ps}
}

func (us *usersService) GetById(id int) *dto.UserDto {
	model := us.repository.GetById(id)
	if model == nil {
		return nil
	}
	// TODO: add automapper
	return &dto.UserDto{
		Id:    model.Id,
		Name:  model.Name,
		Email: model.Email,
	}
}

func (us *usersService) Create(obj *dto.CreateUserDto) *dto.UserDto {
	// TODO: add automapper
	passwordModel := us.passwordService.CreatePasswordAndGenerateSalt(obj.Password)

	v := us.passwordService.VerifyPassword(passwordModel, passwordModel.Hash)
	if !v {
		panic("NOT VALID")
	}

	model := &users.UserModel{
		Name:  obj.Name,
		Email: obj.Email,
		Salt:  passwordModel.Salt,
		Hash:  passwordModel.Hash,
	}

	model = us.repository.Create(model)
	return &dto.UserDto{
		Id:    model.Id,
		Name:  model.Name,
		Email: model.Email,
	}
}
