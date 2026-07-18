package model

import "time"

type ReportType struct {
	Code        string `json:"code" db:"code"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	IsActive    bool   `json:"is_active" db:"is_active"`
}

type AIReport struct {
	ID             string    `json:"id" db:"id"`
	BirthProfileID string    `json:"birth_profile_id" db:"birth_profile_id"`
	ReportType     string    `json:"report_type" db:"report_type"`
	PromptVersion  *string   `json:"prompt_version,omitempty" db:"prompt_version"`
	Model          *string   `json:"model,omitempty" db:"model"`
	Content        *string   `json:"content,omitempty" db:"content"`
	Language       string    `json:"language" db:"language"`
	Status         string    `json:"status" db:"status"`
	GeneratedBy    *string   `json:"generated_by,omitempty" db:"generated_by"`
	Version        int       `json:"version" db:"version"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
