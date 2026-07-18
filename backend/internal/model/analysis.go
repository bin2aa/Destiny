package model

import "time"

type AnalysisResult struct {
	ID             string    `json:"id" db:"id"`
	BirthProfileID string    `json:"birth_profile_id" db:"birth_profile_id"`
	AnalysisType   string    `json:"analysis_type" db:"analysis_type"`
	Data           []byte    `json:"data" db:"data"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
