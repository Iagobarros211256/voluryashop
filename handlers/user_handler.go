package handlers

import (
	"net/http"

	"github.com/Iagobarros211256/voluryashop/services"
)


type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(W, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

///////
package handlers

import (
    "encoding/json"
    "net/http"
    "myapp/models"
    "myapp/services"
)

type OrderHandler struct {
    Service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
    return &OrderHandler{Service: service}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var payload struct {
        UserID int                `json:"user_id"`
        Items  []models.OrderItem `json:"items"`
    }

    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    order, err := h.Service.CreateOrder(r.Context(), payload.UserID, payload.Items)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(order)
}

