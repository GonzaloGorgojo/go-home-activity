package users

import "github.com/gonzalogorgojo/go-home-activity/internal/models"

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetOneByEmail(email string) (models.User, error)
	CreateUser(newUser models.CreateUserRequest) (models.User, error)
}
