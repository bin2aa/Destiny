package model

import "time"

type TarotReading struct {
	ID             string    `json:"id" db:"id"`
	BirthProfileID string    `json:"birth_profile_id" db:"birth_profile_id"`
	Question       *string   `json:"question,omitempty" db:"question"`
	Spread         *string   `json:"spread,omitempty" db:"spread"`
	Result         []byte    `json:"result" db:"result"`
	AISummary      *string   `json:"ai_summary,omitempty" db:"ai_summary"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
