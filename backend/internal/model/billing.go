package model

import "time"

type Plan struct {
	ID           string    `json:"id" db:"id"`
	Code         string    `json:"code" db:"code"`
	Name         string    `json:"name" db:"name"`
	Price        float64   `json:"price" db:"price"`
	Currency     string    `json:"currency" db:"currency"`
	DurationDays int       `json:"duration_days" db:"duration_days"`
	Features     []byte    `json:"features,omitempty" db:"features"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type Subscription struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	PlanID    string    `json:"plan_id" db:"plan_id"`
	StartAt   time.Time `json:"start_at" db:"start_at"`
	EndAt     time.Time `json:"end_at" db:"end_at"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Payment struct {
	ID             string     `json:"id" db:"id"`
	SubscriptionID string     `json:"subscription_id" db:"subscription_id"`
	Provider       string     `json:"provider" db:"provider"`
	Amount         float64    `json:"amount" db:"amount"`
	Currency       string     `json:"currency" db:"currency"`
	Status         string     `json:"status" db:"status"`
	TransactionID  *string    `json:"transaction_id,omitempty" db:"transaction_id"`
	PaidAt         *time.Time `json:"paid_at,omitempty" db:"paid_at"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}
