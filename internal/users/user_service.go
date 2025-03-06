package users

import (
	"github.com/gonzalogorgojo/go-home-activity/internal/models"
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
