package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gonzalogorgojo/go-home-activity/internal/config"
	"github.com/gonzalogorgojo/go-home-activity/internal/models"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(user *models.User, expirationTime time.Duration) (string, error) {

	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expirationTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    cfg.JWTIssuer,
			Subject:   user.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTToken(tokenString string) (*Claims, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidToken
			}
			return []byte(cfg.JWTSecretKey), nil
		},
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
