package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/model"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(ctx context.Context, order *model.Order) error {
	query := `
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

	_, err := r.db.Exec(
		ctx,
		query,
		order.ID,
		order.UserID,
		order.Status,
		order.TotalAmount,
		order.CreatedAt,
		order.UpdatedAt,
	)

	return err
}

func (r *OrderRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
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
