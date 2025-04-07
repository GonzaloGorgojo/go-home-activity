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
	err := r.DB.QueryRow("SELECT ID, Password, Status FROM User WHERE Email = ?", email).Scan(&u.ID, &u.Password, &u.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if u.Status == utils.Suspended {
		log.Printf("User %v is suspended", u.ID)

		return nil, utils.ErrSuspendedUser
	}

	return u, nil
}

func (r *AuthRepositoryImpl) checkValidUser(id int64) error {
	u := &models.User{}

	err := r.DB.QueryRow("SELECT Status FROM User WHERE ID = ?", id).Scan(&u.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.ErrUserNotFound
		}
		return err
	}
	if u.Status == utils.Suspended {
		log.Printf("User %v is suspended", id)
		return utils.ErrSuspendedUser
	}

	return nil
}

func (r *AuthRepositoryImpl) LogIn(req models.LogInRequest) (*models.User, error) {
	return r.getUserByEmail(req.Email)
}

func (r *AuthRepositoryImpl) SignUp(req models.SignUpRequest) (*int64, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}

	existingUser, err := r.getUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, utils.ErrEmailAlreadyInUse
	}

	result, err := tx.Exec("INSERT INTO User (name, email, password, type) VALUES (?, ?, ?, ?)",
		req.Name, req.Email, req.Password, "free")

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	refreshToken, err := utils.GenerateJWTToken(id, utils.RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("INSERT INTO UserToken (userID, refreshToken) VALUES (?, ?)",
		id, refreshToken)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *AuthRepositoryImpl) SearchRefreshToken(id int64) (*string, error) {
	var refreshToken string

	err := r.checkValidUser(id)
	if err != nil {
		return nil, err
	}

	err = r.DB.QueryRow("SELECT refreshToken FROM UserToken WHERE userID = ?", id).Scan(&refreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &refreshToken, nil
}

func (r *AuthRepositoryImpl) UpdateRefreshToken(id int64, token string) error {
	result, err := r.DB.Exec(`
		UPDATE UserToken 
		SET refreshToken = ?, updatedAt = strftime('%Y-%m-%d %H:%M:%f', 'now')
		WHERE userID = ?`,
		token, id)

	if err != nil {
		return fmt.Errorf("error updating refresh token: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return utils.ErrNoRowsUpdated
	}

	log.Printf("Refresh Token was updated for user %v", id)
	return nil
}
