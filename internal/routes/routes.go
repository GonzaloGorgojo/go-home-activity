package routes

import (
	"database/sql"
	"net/http"

	"github.com/gonzalogorgojo/go-home-activity/internal/handlers"
)

func AddRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.Handle("GET /users", http.HandlerFunc(handlers.GetAllUsers))
}
