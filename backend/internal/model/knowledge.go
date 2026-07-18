package model

import "time"

type KnowledgeBase struct {
	ID             string    `json:"id" db:"id"`
	Category       *string   `json:"category,omitempty" db:"category"`
	Title          string    `json:"title" db:"title"`
	Content        string    `json:"content" db:"content"`
	EmbeddingModel *string   `json:"embedding_model,omitempty" db:"embedding_model"`
	Language       string    `json:"language" db:"language"`
	Source         *string   `json:"source,omitempty" db:"source"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type PromptTemplate struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	SystemPrompt string    `json:"system_prompt" db:"system_prompt"`
	Version      int       `json:"version" db:"version"`
	Active       bool      `json:"active" db:"active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
