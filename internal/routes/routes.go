package routes

import (
	"net/http"

	"github.com/gonzalogorgojo/go-home-activity/internal/auth"
	"github.com/gonzalogorgojo/go-home-activity/internal/middleware"
	"github.com/gonzalogorgojo/go-home-activity/internal/users"
)

func AddRoutes(mux *http.ServeMux, userHandler *users.UserHandler, authhandler *auth.AuthHandler) {

	mux.HandleFunc("POST /login", authhandler.Login)
	mux.HandleFunc("POST /signup", authhandler.SignUp)

	mux.Handle("GET /users", middleware.AuthMiddleware(http.HandlerFunc(userHandler.GetAllUsers)))
}
