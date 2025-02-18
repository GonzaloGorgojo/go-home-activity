package routes

import (
	"database/sql"
	"net/http"

	"github.com/gonzalogorgojo/go-home-activity/internal/handlers"
	"github.com/gonzalogorgojo/go-home-activity/internal/repositories"
	"github.com/gonzalogorgojo/go-home-activity/internal/services"
)

func AddRoutes(mux *http.ServeMux, db *sql.DB) {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	mux.HandleFunc("GET /users", handler.GetAllUsers)
}
