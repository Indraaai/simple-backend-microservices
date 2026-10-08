package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order, items []model.OrderItem) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Order, error)
}

type orderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Create(ctx context.Context, order *model.Order, items []model.OrderItem) error {
	if r == nil || r.db == nil {
		return errors.New("order repository database is nil")
	}
	if order == nil {
		return errors.New("order is nil")
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	orderQuery := `
		INSERT INTO orders (
			id,
			user_id,
			status,
			total_amount,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = tx.Exec(
		ctx,
		orderQuery,
		order.ID,
		order.UserID,
		order.Status,
		order.TotalAmount,
		order.CreatedAt,
		order.UpdatedAt,
	)

	if err != nil {
		return err
	}

	itemQuery := `
		INSERT INTO order_items (
			id,
			order_id,
			product_id,
			quantity,
			price,
			subtotal
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	for _, item := range items {
		_, err = tx.Exec(
			ctx,
			itemQuery,
			item.ID,
			item.OrderID,
			item.ProductID,
			item.Quantity,
			item.Price,
			item.SubTotal,
		)

		if err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	if r == nil || r.db == nil {
		return nil, errors.New("order repository database is nil")
	}

	query := `
		SELECT
			id,
			user_id,
			status,
			total_amount,
			created_at,
			updated_at
		FROM orders
		WHERE id = $1
	`

	order := &model.Order{}

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&order.ID,
		&order.UserID,
		&order.Status,
		&order.TotalAmount,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}
