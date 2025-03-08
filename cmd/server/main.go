package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gonzalogorgojo/go-home-activity/internal/auth"
	"github.com/gonzalogorgojo/go-home-activity/internal/config"
	"github.com/gonzalogorgojo/go-home-activity/internal/database"
	"github.com/gonzalogorgojo/go-home-activity/internal/routes"
	"github.com/gonzalogorgojo/go-home-activity/internal/users"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.InitDB(cfg.DBConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	userRepo := users.NewUserRepositoryImpl(db)
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserHandler(userService)

	authRepo := auth.NewAuthRepositoryImpl(db)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	routes.AddRoutes(mux, userHandler, authHandler)

	s := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting server on port %s...", cfg.Port)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
