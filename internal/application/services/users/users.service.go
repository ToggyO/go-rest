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
	mapper          contracts.IMapper
}

func NewUsersService(r repositories.IUsersRepository,
	ps contracts.IPasswordService, m contracts.IMapper) IUsersService {
	return &usersService{r, ps, m}
}

func (us *usersService) GetById(id int) *dto.UserDto {
	model := us.repository.GetById(id)
	if model == nil {
		return nil
	}

	userDto := new(dto.UserDto)
	us.mapper.Map(model, userDto)
	return userDto
}

func (us *usersService) Create(obj *dto.CreateUserDto) *dto.UserDto {

	passwordModel := us.passwordService.CreatePasswordAndGenerateSalt(obj.Password)
	model := &users.UserModel{
		Name:  obj.Name,
		Email: obj.Email,
		Salt:  passwordModel.Salt,
		Hash:  passwordModel.Hash,
	}

	model = us.repository.Create(model)

	userDto := new(dto.UserDto)
	us.mapper.Map(model, userDto)
	return userDto
}

func (us *usersService) Update(obj *dto.UpdateUserDto) *dto.UserDto {
	model := us.repository.GetById(obj.Id)
	if model == nil {
		return nil
	}

	us.mapper.Map(obj, model)
	us.repository.Update(model)

	userDto := new(dto.UserDto)
	us.mapper.Map(model, userDto)
	return userDto
}

func (us *usersService) Delete(id int) {
	us.repository.Delete(id)
}
