package auth

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gonzalogorgojo/go-home-activity/internal/models"
	"github.com/gonzalogorgojo/go-home-activity/internal/utils"
)

type AuthRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepositoryImpl(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{DB: db}
}

func (r *AuthRepositoryImpl) getUserByEmail(email string) (*models.User, error) {
	u := &models.User{}
	err := r.DB.QueryRow("SELECT ID, Name, Email, Password, Type  FROM User WHERE Email = ?", email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Type)
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

func (r *AuthRepositoryImpl) SignUp(req models.SignUpRequest, token string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	existingUser, err := r.getUserByEmail(req.Email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return utils.ErrEmailAlreadyInUse
	}

	_, err = r.DB.Exec("INSERT INTO User (name, email, password, type) VALUES (?, ?, ?, ?)",
		req.Name, req.Email, req.Password, "free")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.DB.Exec("INSERT INTO UserToken (userEmail, refreshToken) VALUES (?, ?)",
		req.Email, token)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *AuthRepositoryImpl) SearchRefreshToken(email string) (*string, error) {
	var refreshToken string

	err := r.DB.QueryRow("SELECT refreshToken FROM UserToken WHERE userEmail = ?", email).Scan(&refreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &refreshToken, nil
}

func (r *AuthRepositoryImpl) SaveRefreshToken(email string, token string) error {
	result, err := r.DB.Exec(`
		UPDATE UserToken 
		SET refreshToken = ?, updatedAt = strftime('%Y-%m-%d %H:%M:%f', 'now')
		WHERE userEmail = ?`,
		token, email)

	if err != nil {
		return fmt.Errorf("error updating refresh token: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated for email: %s", email)
	}

	log.Printf("Refresh Token was updated for user %v", email)
	return nil
}
