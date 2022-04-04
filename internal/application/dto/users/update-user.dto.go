package dto

type UpdateUserDto struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}
