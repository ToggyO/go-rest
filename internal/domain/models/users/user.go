package users

type User struct {
	Id    int
	Name  string
	Email string
	Hash  string
	Salt  string
}
