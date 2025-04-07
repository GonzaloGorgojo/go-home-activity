package auth

import (
	"errors"

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

	shotLivedToken, err := utils.GenerateJWTToken(existingUser.ID, utils.ShortTokenExpiry)
	if err != nil {
		return nil, err
	}

	longLivedToken, err := utils.GenerateJWTToken(existingUser.ID, utils.RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}

	err = s.authRepo.UpdateRefreshToken(existingUser.ID, longLivedToken)

	if err != nil {
		return nil, err
	}

	return &models.LogInResponse{
		Token: shotLivedToken,
	}, nil
}

func (s *AuthService) SignUp(req models.SignUpRequest) (*models.SignUpResponse, error) {

	hashedPass, err := utils.GenerateHashFromPassword(req.Password)
	if err != nil {
		return nil, err
	}
	req.Password = hashedPass

	idPtr, err := s.authRepo.SignUp(req)
	if err != nil {
		return nil, err
	}
	if idPtr == nil {
		return nil, errors.New("id pointer is nil")
	}
	id := *idPtr

	shotLivedToken, err := utils.GenerateJWTToken(id, utils.ShortTokenExpiry)
	if err != nil {
		return nil, err
	}

	return &models.SignUpResponse{
		Token: shotLivedToken,
	}, nil
}

func (s *AuthService) SearchRefreshToken(id int64) (*string, error) {
	storedHash, err := s.authRepo.SearchRefreshToken(id)
	if err != nil || storedHash == nil {
		return nil, errors.New("refresh token not found")
	}

	return storedHash, nil
}

func (s *AuthService) UpdateRefreshToken(id int64, token string) error {
	err := s.authRepo.UpdateRefreshToken(id, token)
	if err != nil {
		return err
	}

	return nil
}
