package dto

import "github.com/google/uuid"

type CreateOrderRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
}
