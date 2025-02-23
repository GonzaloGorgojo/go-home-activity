package routes

import (
	"net/http"

	"github.com/gonzalogorgojo/go-home-activity/internal/handlers"
)

func AddRoutes(mux *http.ServeMux, handler *handlers.UserHandler) {

	mux.HandleFunc("GET /users", handler.GetAllUsers)
}
