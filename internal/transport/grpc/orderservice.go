package grpc

import (
	"context"
	"github.com/AlekSi/pointer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"orders-microservice/internal/models"
	client "orders-microservice/pkg/api/order"
)

type Service interface {
	CreateOrder(ctx context.Context, order models.Order) (*models.Order, error)
	GetOrder(ctx context.Context, order models.Order) (*models.Order, error)
	UpdateOrder(ctx context.Context, order models.Order) (*models.Order, error)
	DeleteOrder(ctx context.Context, order models.Order) (bool, error)
	ListOrders(ctx context.Context) ([]*models.Order, error)
}

type OrderService struct {
	client.UnimplementedOrderServiceServer
	service Service
}

func NewOrderService(srv Service) *OrderService {
	return &OrderService{service: srv}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *client.CreateOrderRequest) (*client.CreateOrderResponse, error) {
	resp, err := s.service.CreateOrder(ctx, models.Order{
		Item:     req.GetItem(),
		Quantity: req.GetQuantity(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "CreateOrder: %s", err)
	}
	r := pointer.Get(resp)
	return &client.CreateOrderResponse{
		Id: r.ID,
	}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *client.GetOrderRequest) (*client.GetOrderResponse, error) {
	resp, err := s.service.GetOrder(ctx, models.Order{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "GetOrder: %s", err)
	}
	r := pointer.Get(resp)
	return &client.GetOrderResponse{
		Order: &client.Order{
			Id:       r.ID,
			Item:     r.Item,
			Quantity: r.Quantity,
		},
	}, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *client.UpdateOrderRequest) (*client.UpdateOrderResponse, error) {
	resp, err := s.service.UpdateOrder(ctx, models.Order{
		ID:       req.GetId(),
		Item:     req.GetItem(),
		Quantity: req.GetQuantity(),
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "UpdateOrder: %s", err)
	}
	r := pointer.Get(resp)
	return &client.UpdateOrderResponse{
		Order: &client.Order{
			Id:       r.ID,
			Item:     r.Item,
			Quantity: r.Quantity,
		},
	}, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, req *client.DeleteOrderRequest) (*client.DeleteOrderResponse, error) {
	resp, err := s.service.DeleteOrder(ctx, models.Order{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "DeleteOrder: %s", err)
	}

	return &client.DeleteOrderResponse{
		Success: resp,
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, req *client.ListOrdersRequest) (*client.ListOrdersResponse, error) {
	resp, err := s.service.ListOrders(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "ListOrders: %s", err)
	}

	result := make([]*client.Order, 0)
	for _, order := range resp {
		r := pointer.Get(order)
		result = append(result, &client.Order{
			Id:       r.ID,
			Item:     r.Item,
			Quantity: r.Quantity,
		})
	}

	return &client.ListOrdersResponse{
		Orders: result,
	}, nil
}
