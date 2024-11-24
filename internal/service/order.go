package service

import (
	"context"
	"orders-microservice/internal/models"
)

type OrderRepo interface {
	CreateOrder(ctx context.Context, order models.Order) (*models.Order, error)
	GetOrder(ctx context.Context, order models.Order) (*models.Order, error)
	UpdateOrder(ctx context.Context, order models.Order) (*models.Order, error)
	DeleteOrder(ctx context.Context, order models.Order) (bool, error)
	ListOrders(ctx context.Context) ([]*models.Order, error)
}

type OrderService struct {
	Repo OrderRepo
}

func NewOrderService(repo OrderRepo) *OrderService {
	return &OrderService{repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	return s.Repo.CreateOrder(ctx, order)
}

func (s *OrderService) GetOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	return s.Repo.GetOrder(ctx, order)
}

func (s *OrderService) UpdateOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	return s.Repo.UpdateOrder(ctx, order)
}

func (s *OrderService) DeleteOrder(ctx context.Context, order models.Order) (bool, error) {
	return s.Repo.DeleteOrder(ctx, order)
}

func (s *OrderService) ListOrders(ctx context.Context) ([]*models.Order, error) {
	return s.Repo.ListOrders(ctx)
}
