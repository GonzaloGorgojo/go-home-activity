package models

type SignUpRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SignUpResponse struct {
	Token string `json:"token"`
}

type LogInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LogInResponse struct {
	Token string `json:"token"`
}
