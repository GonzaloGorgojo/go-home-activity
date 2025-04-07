package models

import (
	"database/sql"
)

type User struct {
	ID        int64        `json:"-"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Password  string       `json:"-"`
	Type      string       `json:"type"`
	Status    string       `json:"status"`
	CreatedAt string       `json:"-"`
	UpdatedAt sql.NullTime `json:"-"`
}

type UserToken struct {
	ID           int64        `json:"id"`
	RefreshToken string       `json:"token"`
	UserID       string       `json:"userID"`
	CreatedAt    string       `json:"-"`
	UpdatedAt    sql.NullTime `json:"-"`
}
