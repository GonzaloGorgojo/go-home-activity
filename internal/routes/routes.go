package routes

import (
	"net/http"

	"github.com/gonzalogorgojo/go-home-activity/internal/auth"
	"github.com/gonzalogorgojo/go-home-activity/internal/middleware"
	"github.com/gonzalogorgojo/go-home-activity/internal/users"
)

func AddRoutes(mux *http.ServeMux, userHandler *users.UserHandler, authhandler *auth.AuthHandler) {

	mux.HandleFunc("POST /login", authhandler.Login)

	mux.Handle("GET /users", middleware.AuthMiddleware(http.HandlerFunc(userHandler.GetAllUsers)))
	mux.Handle("GET /user", middleware.AuthMiddleware(http.HandlerFunc(userHandler.GetOneByEmail)))
	mux.Handle("POST /user", middleware.AuthMiddleware(http.HandlerFunc(userHandler.CreateUser)))
}
