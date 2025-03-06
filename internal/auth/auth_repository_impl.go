package auth

import (
	"database/sql"

	"github.com/gonzalogorgojo/go-home-activity/internal/models"
	"github.com/gonzalogorgojo/go-home-activity/internal/utils"
)

type AuthRepositoryImpl struct {
	Db *sql.DB
}

func NewAuthRepositoryImpl(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{Db: db}
}

func (r *AuthRepositoryImpl) getUserByEmail(email string) (*models.User, error) {
	u := &models.User{}
	err := r.Db.QueryRow("SELECT * FROM users WHERE Email = ?", email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return u, err
}

func (r *AuthRepositoryImpl) LogIn(req models.LogInRequest) (*models.User, error) {
	return r.getUserByEmail(req.Email)

}

func (r *AuthRepositoryImpl) SignUp(req models.SignUpRequest) (*models.User, error) {
	existingUser, err := r.getUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, utils.ErrEmailAlreadyInUse
	}

	_, err = r.Db.Exec("INSERT INTO users (name, email, password, type) VALUES (?, ?, ?, ?)",
		req.Name, req.Email, req.Password, "free")
	if err != nil {
		return nil, err
	}

	newUser, err := r.getUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	return newUser, nil

}
