package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gonzalogorgojo/go-home-activity/internal/utils"
)

type userContextKey struct{}

type AuthMiddleware struct {
	authService *AuthService
}

func NewAuthMiddleware(authService *AuthService) *AuthMiddleware {
	return &AuthMiddleware{authService: authService}
}

func (m *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateJWTToken(tokenString)

		if err != nil {
			if err == utils.ErrExpiredToken {
				refreshToken, err := m.authService.SearchRefreshToken(claims.Email)
				if err != nil || refreshToken == nil {
					http.Error(w, "Refresh token not found", http.StatusUnauthorized)
					return
				}
				log.Printf("Refresh Token was found for user %v", claims.Email)

				claims, err = utils.ValidateJWTToken(*refreshToken)
				if err != nil {
					if err == utils.ErrExpiredToken {
						http.Error(w, "Refresh token was also expired", http.StatusUnauthorized)
					} else {
						http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
					}
					return
				}
				longNewToken, err := utils.GenerateJWTToken(claims.Email, utils.RefreshTokenExpiry)
				if err != nil {
					http.Error(w, "Failed to generate new long lived token", http.StatusInternalServerError)
					return
				}
				err = m.authService.SaveRefreshToken(claims.Email, longNewToken)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				shortNewToken, err := utils.GenerateJWTToken(claims.Email, utils.ShortTokenExpiry)
				if err != nil {
					http.Error(w, "Failed to generate new short lived token", http.StatusInternalServerError)
					return
				}
				r.Header.Set("Authorization", "Bearer "+shortNewToken)
				w.Header().Set("Authorization", "Bearer "+shortNewToken)
			} else {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
		}

		ctx := context.WithValue(r.Context(), userContextKey{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetClaimsFromContext(ctx context.Context) (*utils.Claims, bool) {
	claims, ok := ctx.Value(userContextKey{}).(*utils.Claims)
	return claims, ok
}
