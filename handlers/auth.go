package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"myapp/services"

	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	Secret      string
	AuthService *services.AuthService
	// aqui poderia ter tbm userRepo (para buscar usuários no banco)
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

	// aqui salvaria no banco: payload.Email + hashed
	// mas por simplicidade vamos só devolver
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"email":    payload.Email,
		"password": hashed, // só para teste, em prod não devolve hash
	})
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

	// ⚠️ Simulação → em prod, busca o hash no banco
	savedHash, _ := h.AuthService.HashPassword("123456")

	if !h.AuthService.CheckPasswordHash(payload.Password, savedHash) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	// cria JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": payload.Email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(h.Secret))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
