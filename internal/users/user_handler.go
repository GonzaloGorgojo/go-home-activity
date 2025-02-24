package users

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		log.Printf("Error fetching users from service: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("Error encoding users response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// TODO: remove because its not needed
func (h *UserHandler) GetOneByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if email == "" {
		http.Error(w, "Email parameter is required", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetOneByEmail(email)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User with email %s not found.\n", email)
			http.Error(w, "User not found", http.StatusBadRequest)
			return
		}
		log.Printf("Error fetching user from service: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Error encoding users response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
