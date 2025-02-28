package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gonzalogorgojo/go-home-activity/internal/models"
	"github.com/gonzalogorgojo/go-home-activity/internal/users"
	"github.com/gonzalogorgojo/go-home-activity/internal/utils"
	"github.com/gonzalogorgojo/go-home-activity/internal/validation"
)

type AuthHandler struct {
	service  *AuthService
	userRepo users.UserRepository
}

func NewAuthHandler(service *AuthService, userRepo users.UserRepository) *AuthHandler {
	return &AuthHandler{service: service, userRepo: userRepo}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Invalid JSON Request: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := validation.Validate.Struct(req); err != nil {
		log.Printf("Error with request: %v", err)
		http.Error(w, validation.FormatValidationErrors(err), http.StatusBadRequest)
		return
	}

	existingUser, err := h.userRepo.GetOneByEmail(req.Email)
	if existingUser.ID == 0 {
		log.Printf("User not found: %v", req.Email)
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	encodedHash, err := utils.ComparePasswordAndHash(req.Password, existingUser.Password)
	if err != nil || !encodedHash {
		log.Printf("Invalid password for user: %v", req.Email)
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	loginResponse, err := h.service.Login(req)
	if err != nil {
		log.Printf("Error during login: %v", err)
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loginResponse)
	log.Printf("User %v logged correctly.", req.Email)
}
