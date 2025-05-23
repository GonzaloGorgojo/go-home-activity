package utils

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrTokenGeneration   = errors.New("failed to generate token")
	ErrInternalServer    = errors.New("internal server error")
	ErrEmailAlreadyInUse = errors.New("email already in use")
	ErrInvalidToken      = errors.New("invalid token")
	ErrExpiredToken      = errors.New("token has expired")
	ErrSuspendedUser     = errors.New("user is suspended")
	ErrNoRowsUpdated     = errors.New("no rows updated for user")
)
