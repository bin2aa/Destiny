package model

import "time"

type DailyHoroscope struct {
	ID             string    `json:"id" db:"id"`
	BirthProfileID string    `json:"birth_profile_id" db:"birth_profile_id"`
	Date           string    `json:"date" db:"date"`
	Summary        *string   `json:"summary,omitempty" db:"summary"`
	Career         *string   `json:"career,omitempty" db:"career"`
	Love           *string   `json:"love,omitempty" db:"love"`
	Health         *string   `json:"health,omitempty" db:"health"`
	Finance        *string   `json:"finance,omitempty" db:"finance"`
	LuckyColor     *string   `json:"lucky_color,omitempty" db:"lucky_color"`
	LuckyNumber    *int16    `json:"lucky_number,omitempty" db:"lucky_number"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
