package auth

import "github.com/gonzalogorgojo/go-home-activity/internal/models"

type AuthRepository interface {
	LogIn(req models.LogInRequest) (*models.User, error)
	SignUp(req models.SignUpRequest, token string) error
	SearchRefreshToken(email string) (*string, error)
	SaveRefreshToken(email string, token string) error
}
