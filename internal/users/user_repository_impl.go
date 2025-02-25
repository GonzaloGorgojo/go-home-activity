package users

import (
	"database/sql"

	"github.com/gonzalogorgojo/go-home-activity/internal/models"
)

type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{Db: db}
}

func (r *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	rows, err := r.Db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Type); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepositoryImpl) GetOneByEmail(email string) (models.User, error) {
	var u models.User
	err := r.Db.QueryRow("SELECT * FROM users WHERE Email = ?", email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Type)

	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (r *UserRepositoryImpl) CreateUser(user models.CreateUserRequest) (models.User, error) {
	result, err := r.Db.Exec("INSERT INTO users (name, email, password, type) VALUES (?, ?, ?, ?)",
		user.Name, user.Email, user.Password, user.Type)
	if err != nil {
		return models.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:    int(id),
		Name:  user.Name,
		Email: user.Email,
		Type:  user.Type,
	}, nil
}
