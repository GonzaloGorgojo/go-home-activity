package repositories

import "github.com/gonzalogorgojo/go-home-activity/internal/models"

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
}
