package dto

type UpdateUserRequest struct {
	Name     *string `json:"name" validate:"omitempty,min=2,max=150"`
	Avatar   *string `json:"avatar"`
	Language *string `json:"language" validate:"omitempty,min=2,max=10"`
	Timezone *string `json:"timezone" validate:"omitempty,max=50"`
	Country  *string `json:"country" validate:"omitempty,max=2"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6,max=100"`
}

type UserResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Avatar       *string `json:"avatar,omitempty"`
	Language     string  `json:"language"`
	Timezone     string  `json:"timezone"`
	Country      *string `json:"country,omitempty"`
	IsPremium    bool    `json:"is_premium"`
	PremiumUntil *string `json:"premium_until,omitempty"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
