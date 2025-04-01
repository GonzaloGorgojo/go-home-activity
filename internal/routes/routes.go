package routes

import (
	"net/http"

	"github.com/gonzalogorgojo/go-home-activity/internal/auth"
	"github.com/gonzalogorgojo/go-home-activity/internal/users"
)

func AddRoutes(mux *http.ServeMux, userHandler *users.UserHandler, authhandler *auth.AuthHandler, authMiddleware *auth.AuthMiddleware) {

	mux.HandleFunc("POST /login", authhandler.Login)
	mux.HandleFunc("POST /signup", authhandler.SignUp)

	mux.Handle("/users", authMiddleware.AuthMiddleware(http.HandlerFunc(userHandler.GetAllUsers)))
}
