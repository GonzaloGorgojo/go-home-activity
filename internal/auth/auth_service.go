package auth

import (
	"github.com/gonzalogorgojo/go-home-activity/internal/models"
	"github.com/gonzalogorgojo/go-home-activity/internal/utils"

	"github.com/gonzalogorgojo/go-home-activity/internal/users"
)

type AuthService struct {
	userRepo users.UserRepository
}

func NewAuthService(userRepo users.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(req models.LoginUserRequest) (*models.LoginResponse, error) {

	existingUser, err := s.userRepo.GetOneByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser.ID == 0 {
		return nil, ErrUserNotFound
	}

	validPassword, err := utils.ComparePasswordAndHash(req.Password, existingUser.Password)
	if err != nil || !validPassword {
		return nil, ErrInvalidPassword
	}

	token, err := utils.GenerateToken(&existingUser)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
	}, nil
}
