package users

type UserModel struct {
	Id    int
	Name  string
	Email string
	Hash  string
	Salt  string
}
