package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mock service
type mockUserService struct{}

func (m *mockUserService) ListUsers(ctx context.Context) ([]models.User, error) {
	return []models.User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
	}, nil
}

func TestUsersHandler(t *testing.T) {
	service := &mockUserService{}
	handler := NewUserHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	handler.UsersHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(users) != 1 || users[0].Name != "Alice" {
		t.Fatalf("unexpected response: %+v", users)
	}
}
