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
	// TODO: add automapper
	var dto *dto.UserDto
	us.mapper.Map(model, dto)
	return dto
	//return &dto.UserDto{
	//	Id:    model.Id,
	//	Name:  model.Name,
	//	Email: model.Email,
	//}
}

func (us *usersService) Create(obj *dto.CreateUserDto) *dto.UserDto {
	// TODO: add automapper
	passwordModel := us.passwordService.CreatePasswordAndGenerateSalt(obj.Password)
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
