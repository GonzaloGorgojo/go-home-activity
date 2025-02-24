package users

type UserRepository interface {
	GetAllUsers() ([]User, error)
	GetOneByEmail(email string) (User, error)
}
