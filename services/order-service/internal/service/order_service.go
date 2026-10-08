package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/dto"
	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/model"
	"github.com/Indraaai/backend-microservices-lab/services/order-service/internal/repository"
)

type OrderService struct {
	repository *repository.OrderRepository
}

func NewOrderService(repository *repository.OrderRepository) *OrderService {
	return &OrderService{
		repository: repository,
	}
}

func (s *OrderService) Create(ctx context.Context, req dto.CreateOrderRequest) (*model.Order, error) {

	now := time.Now()

	order := &model.Order{
		ID:          uuid.New(),
		UserID:      req.UserID,
		Status:      "pending",
		TotalAmount: 0,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repository.Create(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}
