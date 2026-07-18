package model

import "time"

type User struct {
	ID           string     `json:"id" db:"id"`
	Name         string     `json:"name" db:"name"`
	Email        string     `json:"email" db:"email"`
	PasswordHash string     `json:"-" db:"password_hash"`
	Avatar       *string    `json:"avatar,omitempty" db:"avatar"`
	Language     string     `json:"language" db:"language"`
	Timezone     string     `json:"timezone" db:"timezone"`
	Country      *string    `json:"country,omitempty" db:"country"`
	PremiumUntil *time.Time `json:"premium_until,omitempty" db:"premium_until"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}
