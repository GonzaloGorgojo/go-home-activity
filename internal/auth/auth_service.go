package auth

import (
	"github.com/gonzalogorgojo/go-home-activity/internal/models"
)

type AuthService struct {
	authRepo AuthRepository
}

func NewAuthService(authRepo AuthRepository) *AuthService {
	return &AuthService{authRepo: authRepo}
}

func (s *AuthService) Login(loginReq models.LoginUserRequest) (models.LoginResponse, error) {
	loginResponse, err := s.authRepo.Login(loginReq)
	if err != nil {
		return models.LoginResponse{}, err
	}
	return loginResponse, nil
}
