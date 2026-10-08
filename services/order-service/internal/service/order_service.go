package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/dto"
	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/model"
	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/repository"
)

type OrderService struct {
	repository repository.OrderRepository
}

func NewOrderService(repository repository.OrderRepository) *OrderService {
	return &OrderService{
		repository: repository,
	}
}

func (s *OrderService) Create(
	ctx context.Context,
	req dto.CreateOrderRequest,
) (*model.Order, error) {
	if s == nil || s.repository == nil {
		return nil, errors.New("order service repository is nil")
	}

	now := time.Now()

	order := &model.Order{
		ID:          uuid.New(),
		UserID:      req.UserID,
		Status:      "pending",
		TotalAmount: 0,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	items := make([]model.OrderItem, 0, len(req.Items))

	for _, itemRequest := range req.Items {

		subtotal := float64(itemRequest.Quantity) * itemRequest.Price

		order.TotalAmount += subtotal

		item := model.OrderItem{
			ID:        uuid.New(),
			OrderID:   order.ID,
			ProductID: itemRequest.ProductID,
			Quantity:  itemRequest.Quantity,
			Price:     itemRequest.Price,
			SubTotal:  subtotal,
		}

		items = append(items, item)
	}

	if err := s.repository.Create(
		ctx,
		order,
		items,
	); err != nil {
		return nil, err
	}

	return order, nil
}
func (s *OrderService) GetByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	if s == nil || s.repository == nil {
		return nil, errors.New("order service repository is nil")
	}

	return s.repository.GetByID(ctx, id)
}
