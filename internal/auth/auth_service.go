package auth

import (
	"github.com/gonzalogorgojo/go-home-activity/internal/models"
	"github.com/gonzalogorgojo/go-home-activity/internal/utils"
)

type AuthService struct {
	authRepo AuthRepository
}

func NewAuthService(authRepo AuthRepository) *AuthService {
	return &AuthService{authRepo: authRepo}
}

func (s *AuthService) LogIn(req models.LogInRequest) (*models.LogInResponse, error) {

	existingUser, err := s.authRepo.LogIn(req)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, utils.ErrUserNotFound
	}

	validPassword, err := utils.ComparePasswordAndHash(req.Password, existingUser.Password)
	if err != nil || !validPassword {
		return nil, utils.ErrInvalidPassword
	}

	token, err := utils.GenerateJWTToken(existingUser)
	if err != nil {
		return nil, err
	}

	return &models.LogInResponse{
		Token: token,
	}, nil
}

func (s *AuthService) SignUp(req models.SignUpRequest) (*models.SignUpResponse, error) {
	hashedPass, err := utils.GenerateFromPassword(req.Password)
	if err != nil {
		return nil, err
	}
	req.Password = hashedPass

	newUser, err := s.authRepo.SignUp(req)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWTToken(newUser)
	if err != nil {
		return nil, err
	}

	return &models.SignUpResponse{
		Token: token,
	}, nil
}
