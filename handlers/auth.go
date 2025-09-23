package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Iagobarros211256/voluryashop/services"
	"github.com/golang-jwt/jwt/v5"
)

// creating  auth handler
type AuthHandler struct {
	Secret      string
	AuthService *services.AuthService
	// implement later an userRepo (to search for users on database)
}

func NewAuthHandler(secret string, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{Secret: secret, AuthService: authService}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	hashed, err := h.AuthService.HashPassword(payload.Password)
	if err != nil {
		http.Error(w, "could not hash password", http.StatusInternalServerError)
		return
	}

	// ⚠️ save in database: payload.Email + hashed

}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// ⚠️ searchs for hash on database

	// creates JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": payload.Email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(h.Secret))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
