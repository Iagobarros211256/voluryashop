package services

import (
	"context"

	"github.com/Iagobarros211256/voluryashop/repositories"
)

//user
type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) ListUsers(ctx context.Context) ([]models.User, error) {
	return s.Repo.GetAll(ctx)
}

//orders

package services

import (
    "context"
    "errors"
    "myapp/models"
    "myapp/repositories"
)

type OrderService struct {
    OrderRepo   *repositories.OrderRepository
    ProductRepo *repositories.ProductRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository, productRepo *repositories.ProductRepository) *OrderService {
    return &OrderService{OrderRepo: orderRepo, ProductRepo: productRepo}
}

func (s *OrderService) CreateOrder(ctx context.Context, userID int, items []models.OrderItem) (*models.Order, error) {
    var total float64

    for _, item := range items {
        product, err := s.ProductRepo.GetByID(ctx, item.ProductID)
        if err != nil {
            return nil, err
        }

        if product.Stock < item.Quantity {
            return nil, errors.New("insufficient stock")
        }

        total += float64(item.Quantity) * product.Price
    }

    // Cria o pedido
    order, err := s.OrderRepo.Create(ctx, userID, items, total)
    if err != nil {
        return nil, err
    }

    // Atualiza o estoque
    for _, item := range items {
        if err := s.ProductRepo.DecreaseStock(ctx, item.ProductID, item.Quantity); err != nil {
            return nil, err
        }
    }

    return order, nil
}
