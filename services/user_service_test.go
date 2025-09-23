package services


package services

import (
	"context"
	"myapp/models"
	"testing"
)

// mock repository
type mockUserRepo struct{}

func (m *mockUserRepo) GetAll(ctx context.Context) ([]models.User, error) {
	return []models.User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}, nil
}

func TestListUsers(t *testing.T) {
	repo := &mockUserRepo{}
	service := NewUserService(repo)

	users, err := service.ListUsers(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
}
