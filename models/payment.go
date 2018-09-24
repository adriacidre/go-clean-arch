package models

import (
	"time"
)

// Payment struct representation of a payment resource.
type Payment struct {
	ID           int64     `json:"id"`
	PaymentID    string    `json:"payment_id" validate:"required"`
	Organisation string    `json:"organisation_id" validate:"required"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}
