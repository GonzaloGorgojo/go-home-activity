package users

import (
	"database/sql"

	"github.com/gonzalogorgojo/go-home-activity/internal/models"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT Name, Email, Type FROM User")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Name, &u.Email, &u.Type); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
