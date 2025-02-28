package auth

import "github.com/gonzalogorgojo/go-home-activity/internal/models"

type AuthRepository interface {
	Login(loginReq models.LoginUserRequest) (models.LoginResponse, error)
}
