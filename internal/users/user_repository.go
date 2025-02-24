package users

type UserRepository interface {
	GetAllUsers() ([]User, error)
}
