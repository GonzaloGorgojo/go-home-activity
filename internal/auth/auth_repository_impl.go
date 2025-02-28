package auth

import (
	"database/sql"

	"github.com/gonzalogorgojo/go-home-activity/internal/models"
)

type AuthRepositoryImpl struct {
	Db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{Db: db}
}

func (r *AuthRepositoryImpl) Login(loginReq models.LoginUserRequest) (models.LoginResponse, error) {
	rows, err := r.Db.Query("SELECT * FROM users WHERE Email = ?", loginReq.Email)
	if err != nil {
		return models.LoginResponse{}, err
	}
	defer rows.Close()

	token := "generatedToken"

	return models.LoginResponse{Token: token}, nil

}
