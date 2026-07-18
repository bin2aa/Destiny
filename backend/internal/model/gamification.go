package model

import "time"

type Achievement struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Icon        string `json:"icon" db:"icon"`
}

type UserAchievement struct {
	ID            string    `json:"id" db:"id"`
	UserID        string    `json:"user_id" db:"user_id"`
	AchievementID string    `json:"achievement_id" db:"achievement_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

type Notification struct {
	ID        string     `json:"id" db:"id"`
	UserID    string     `json:"user_id" db:"user_id"`
	Title     string     `json:"title" db:"title"`
	Body      *string    `json:"body,omitempty" db:"body"`
	Type      *string    `json:"type,omitempty" db:"type"`
	ReadAt    *time.Time `json:"read_at,omitempty" db:"read_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
}
