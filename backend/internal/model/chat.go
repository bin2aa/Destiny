package model

import "time"

type AIChat struct {
	ID             string    `json:"id" db:"id"`
	UserID         string    `json:"user_id" db:"user_id"`
	BirthProfileID *string   `json:"birth_profile_id,omitempty" db:"birth_profile_id"`
	Title          *string   `json:"title,omitempty" db:"title"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type AIMessage struct {
	ID        string    `json:"id" db:"id"`
	ChatID    string    `json:"chat_id" db:"chat_id"`
	Role      string    `json:"role" db:"role"`
	Content   string    `json:"content" db:"content"`
	Token     *int      `json:"token,omitempty" db:"token"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
