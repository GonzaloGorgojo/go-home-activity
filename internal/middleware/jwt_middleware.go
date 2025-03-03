package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gonzalogorgojo/go-home-activity/internal/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
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

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			statusCode := http.StatusUnauthorized
			message := "Invalid token"

			if err == utils.ErrExpiredToken {
				message = "Token has expired"
			}

			http.Error(w, message, statusCode)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
