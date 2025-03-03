package auth

type AuthRepository interface {
	SignUp(string) (bool, error)
}
