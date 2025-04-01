package auth

import (
	"errors"
	"log"

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

	shotLivedToken, err := utils.GenerateJWTToken(existingUser.Email, utils.ShortTokenExpiry)
	if err != nil {
		return nil, err
	}

	longLivedToken, err := utils.GenerateJWTToken(existingUser.Email, utils.RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}

	err = s.authRepo.SaveRefreshToken(existingUser.Email, longLivedToken)

	if err != nil {
		return nil, err
	}

	return &models.LogInResponse{
		Token: shotLivedToken,
	}, nil
}

func (s *AuthService) SignUp(req models.SignUpRequest) (*models.SignUpResponse, error) {
	shotLivedToken, err := utils.GenerateJWTToken(req.Email, utils.ShortTokenExpiry)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateJWTToken(req.Email, utils.RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}

	hashedPass, err := utils.GenerateHashFromPassword(req.Password)
	if err != nil {
		return nil, err
	}
	req.Password = hashedPass

	err = s.authRepo.SignUp(req, refreshToken)

	if err != nil {
		return nil, err
	}

	log.Printf("sad %v", refreshToken)

	return &models.SignUpResponse{
		Token: shotLivedToken,
	}, nil
}

func (s *AuthService) SearchRefreshToken(email string) (*string, error) {
	storedHash, err := s.authRepo.SearchRefreshToken(email)
	if err != nil || storedHash == nil {
		return nil, errors.New("refresh token not found")
	}

	return storedHash, nil
}

func (s *AuthService) SaveRefreshToken(email string, token string) error {
	err := s.authRepo.SaveRefreshToken(email, token)
	if err != nil {
		return err
	}

	return nil
}
