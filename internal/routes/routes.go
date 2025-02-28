package routes

import (
	"net/http"

	"github.com/gonzalogorgojo/go-home-activity/internal/auth"
	"github.com/gonzalogorgojo/go-home-activity/internal/users"
)

func AddRoutes(mux *http.ServeMux, userHandler *users.UserHandler, authhandler *auth.AuthHandler) {

	mux.HandleFunc("GET /users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /user", userHandler.GetOneByEmail)
	mux.HandleFunc("POST /user", userHandler.CreateUser)

	mux.HandleFunc("POST /login", authhandler.Login)

}
