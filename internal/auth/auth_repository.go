package auth

import "github.com/gonzalogorgojo/go-home-activity/internal/models"

type AuthRepository interface {
	LogIn(req models.LogInRequest) (*models.User, error)
	SignUp(req models.SignUpRequest) (*int64, error)
	SearchRefreshToken(id int64) (*string, error)
	UpdateRefreshToken(id int64, token string) error
}
