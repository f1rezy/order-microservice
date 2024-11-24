package repository

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"orders-microservice/internal/models"
	"orders-microservice/pkg/db/postgres"
)

type OrderRepository struct {
	db *postgres.DB
}

func NewOrderRepository(db *postgres.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (s *OrderRepository) CreateOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	var result models.Order
	err := sq.Insert("orders").
		Columns("item", "quantity").
		Values(order.Item, order.Quantity).
		Suffix("returning *").
		PlaceholderFormat(sq.Dollar).
		RunWith(s.db.Db).
		QueryRow().
		Scan(&result.ID, &result.Item, &result.Quantity)
	if err != nil {
		return nil, fmt.Errorf("repository.CreateOrder: %w", err)
	}

	return &result, nil
}

func (s *OrderRepository) GetOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	var result models.Order
	err := sq.Select("*").
		From("orders").
		Where(sq.Eq{"id": order.ID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(s.db.Db).
		QueryRow().
		Scan(&result.ID, &result.Item, &result.Quantity)
	if err != nil {
		return nil, fmt.Errorf("repository.GetOrder: %w", err)
	}

	return &result, nil
}

func (s *OrderRepository) UpdateOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	var result models.Order
	err := sq.Update("orders").
		Where(sq.Eq{"id": order.ID}).
		Set("item", order.Item).
		Set("quantity", order.Quantity).
		PlaceholderFormat(sq.Dollar).
		Suffix("returning *").
		RunWith(s.db.Db).
		QueryRow().
		Scan(&result.ID, &result.Item, &result.Quantity)
	if err != nil {
		return nil, fmt.Errorf("repository.UpdateOrder: %w", err)
	}

	return &result, nil
}

func (s *OrderRepository) DeleteOrder(ctx context.Context, order models.Order) (bool, error) {
	result, err := sq.Delete("orders").
		Where(sq.Eq{"id": order.ID}).
		PlaceholderFormat(sq.Dollar).
		Suffix("returning *").
		RunWith(s.db.Db).
		Exec()

	if err != nil {
		return false, fmt.Errorf("repository.DeleteOrder: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("repository.DeleteOrder: %w", err)
	}
	if rowsAffected == 0 {
		return false, fmt.Errorf("repository.DeleteOrder: order (id %s) not found", order.ID)
	}

	return true, nil
}

func (s *OrderRepository) ListOrders(ctx context.Context) ([]*models.Order, error) {
	var result []*models.Order
	rows, err := sq.Select("*").
		From("orders").
		PlaceholderFormat(sq.Dollar).
		RunWith(s.db.Db).
		Query()

	if err != nil {
		return nil, fmt.Errorf("repository.ListOrders: %w", err)
	}

	var id string
	var item string
	var quantity int32

	for rows.Next() {
		err = rows.Scan(&id, &item, &quantity)
		if err != nil {
			return nil, fmt.Errorf("repository.ListOrders: %w", err)
		}
		result = append(result, &models.Order{ID: id, Item: item, Quantity: quantity})
	}

	return result, nil
}
