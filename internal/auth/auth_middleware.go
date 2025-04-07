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
				refreshToken, err := m.authService.SearchRefreshToken(claims.UserID)
				if err != nil {
					log.Printf("Error searching token: %v", err)
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return

				}
				if refreshToken == nil {
					log.Printf("Refresh token not found for user %v", claims.UserID)
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				log.Printf("Refresh Token was found for user %v", claims.UserID)

				claims, err = utils.ValidateJWTToken(*refreshToken)
				if err != nil {
					if err == utils.ErrExpiredToken {
						log.Printf("Refresh token was also expired for user %v", claims.UserID)
						http.Error(w, "Unauthorized", http.StatusUnauthorized)
					} else {
						log.Printf("Invalid refresh token for user %v", claims.UserID)
						http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
					}
					return
				}
				longNewToken, err := utils.GenerateJWTToken(claims.UserID, utils.RefreshTokenExpiry)
				if err != nil {
					log.Printf("Failed to generate new long lived token for user %v", claims.UserID)
					http.Error(w, "Unauthorized", http.StatusInternalServerError)
					return
				}
				err = m.authService.UpdateRefreshToken(claims.UserID, longNewToken)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				log.Printf("Updated long lived token for user %v", claims.UserID)

				shortNewToken, err := utils.GenerateJWTToken(claims.UserID, utils.ShortTokenExpiry)
				if err != nil {
					log.Printf("Failed to generate new short lived token for user %v", claims.UserID)
					http.Error(w, "Unauthorized", http.StatusInternalServerError)
					return
				}
				r.Header.Set("Authorization", "Bearer "+shortNewToken)
				w.Header().Set("Authorization", "Bearer "+shortNewToken)
			} else {
				log.Printf("Failed to generate new short lived token for user %v", claims.UserID)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
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
