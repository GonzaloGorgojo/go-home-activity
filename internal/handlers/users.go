package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gonzalogorgojo/go-home-activity/internal/database"
	"github.com/gonzalogorgojo/go-home-activity/internal/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := database.InitDB()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("Error querying users: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Type); err != nil {
			log.Printf("Error scanning user: %v", err)
			continue
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
