package model

import "time"

type BirthProfile struct {
	ID            string    `json:"id" db:"id"`
	UserID        string    `json:"user_id" db:"user_id"`
	FullName      string    `json:"full_name" db:"full_name"`
	Gender        *string   `json:"gender,omitempty" db:"gender"`
	BirthDate     string    `json:"birth_date" db:"birth_date"`
	BirthTime     *string   `json:"birth_time,omitempty" db:"birth_time"`
	Timezone      string    `json:"timezone" db:"timezone"`
	Latitude      *float64  `json:"latitude,omitempty" db:"latitude"`
	Longitude     *float64  `json:"longitude,omitempty" db:"longitude"`
	City          *string   `json:"city,omitempty" db:"city"`
	Country       *string   `json:"country,omitempty" db:"country"`
	BirthPlace    *string   `json:"birth_place,omitempty" db:"birth_place"`
	IsUnknownTime bool      `json:"is_unknown_time" db:"is_unknown_time"`
	Note          *string   `json:"note,omitempty" db:"note"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
