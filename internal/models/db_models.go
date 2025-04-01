package models

import (
	"database/sql"
)

type User struct {
	ID        uint         `json:"-"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Password  string       `json:"-"`
	Type      string       `json:"type"`
	CreatedAt string       `json:"-"`
	UpdatedAt sql.NullTime `json:"-"`
}

type UserToken struct {
	ID           uint         `json:"id"`
	RefreshToken string       `json:"token"`
	UserEmail    string       `json:"email"`
	CreatedAt    string       `json:"-"`
	UpdatedAt    sql.NullTime `json:"-"`
}
