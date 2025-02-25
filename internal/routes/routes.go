package routes

import (
	"net/http"

	"github.com/gonzalogorgojo/go-home-activity/internal/users"
)

func AddRoutes(mux *http.ServeMux, handler *users.UserHandler) {

	mux.HandleFunc("GET /users", handler.GetAllUsers)
	mux.HandleFunc("GET /user", handler.GetOneByEmail)
	mux.HandleFunc("POST /user", handler.CreateUser)
}
