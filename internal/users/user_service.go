package users

import (
	"errors"

	"github.com/gonzalogorgojo/go-home-activity/internal/models"
	"github.com/gonzalogorgojo/go-home-activity/internal/utils"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetOneByEmail(email string) (models.User, error) {
	return s.repo.GetOneByEmail(email)
}

func (s *UserService) CreateUser(user models.CreateUserRequest) (models.User, error) {
	existingUser, err := s.repo.GetOneByEmail(user.Email)
	if err == nil && existingUser.ID != 0 {
		return models.User{}, errors.New("email already in use")
	}

	encodedHash, err := utils.GenerateFromPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = encodedHash

	return s.repo.CreateUser(user)
}
