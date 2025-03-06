package auth

import "github.com/gonzalogorgojo/go-home-activity/internal/models"

type AuthRepository interface {
	LogIn(req models.LogInRequest) (*models.User, error)
	SignUp(req models.SignUpRequest) (*models.User, error)
}
