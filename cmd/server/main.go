package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gonzalogorgojo/go-home-activity/internal/database"
	"github.com/gonzalogorgojo/go-home-activity/internal/routes"
	"github.com/gonzalogorgojo/go-home-activity/internal/users"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserHandler(userService)

	routes.AddRoutes(mux, userHandler)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server running on port :8080")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
