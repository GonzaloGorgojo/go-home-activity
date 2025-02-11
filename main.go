package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gonzalogorgojo/go-home-activity/internal/database"
	"github.com/gonzalogorgojo/go-home-activity/internal/routes"
)

func main() {

	dbI := database.InitDB()
	defer dbI.Close()

	mux := http.NewServeMux()

	routes.AddRoutes(mux, dbI)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server running on port :8080")
	log.Fatal(s.ListenAndServe())
}
