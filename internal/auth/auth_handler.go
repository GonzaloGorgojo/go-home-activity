package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gonzalogorgojo/go-home-activity/internal/models"
	"github.com/gonzalogorgojo/go-home-activity/internal/utils"
	"github.com/gonzalogorgojo/go-home-activity/internal/validation"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LogInRequest

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

	LogInResponse, err := h.service.LogIn(req)
	if err != nil {
		log.Printf("Error during login: %v", err)
		statusCode := http.StatusInternalServerError
		message := "Internal Server Error"

		if errors.Is(err, utils.ErrUserNotFound) {
			statusCode = http.StatusBadRequest
			message = "User not found"
		} else if errors.Is(err, utils.ErrInvalidPassword) {
			statusCode = http.StatusBadRequest
			message = "Invalid password"
		}

		http.Error(w, message, statusCode)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LogInResponse)
	log.Printf("User %v logged correctly.", req.Email)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req models.SignUpRequest

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

	signUpResponse, err := h.service.SignUp(req)
	if err != nil {
		log.Printf("Error during Signup: %v", err)
		statusCode := http.StatusInternalServerError
		message := "Internal Server Error"

		if errors.Is(err, utils.ErrEmailAlreadyInUse) {
			statusCode = http.StatusBadRequest
			message = "Email already registered"
		}

		http.Error(w, message, statusCode)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(signUpResponse)
	log.Printf("User %v created correctly with refresh token.", req.Email)
}
