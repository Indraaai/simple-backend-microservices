package dto

import "github.com/google/uuid"

type CreateOrderRequest struct {
	UserID uuid.UUID   `json:"user_id" binding:"required"`
	Items  []OrderItem `json:"items" binding:"required,dive"`
}

type OrderItem struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,min=1"`
	Price     float64   `json:"price" binding:"required,gt=0"`
}
