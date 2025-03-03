package auth

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrTokenGeneration = errors.New("failed to generate token")
	ErrInternalServer  = errors.New("internal server error")
)
