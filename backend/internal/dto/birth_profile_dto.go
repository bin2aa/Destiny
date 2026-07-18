package dto

type CreateBirthProfileRequest struct {
	FullName      string   `json:"full_name" validate:"required,min=1,max=150"`
	Gender        *string  `json:"gender"`
	BirthDate     string   `json:"birth_date" validate:"required"`
	BirthTime     *string  `json:"birth_time"`
	Timezone      string   `json:"timezone" validate:"required"`
	Latitude      *float64 `json:"latitude"`
	Longitude     *float64 `json:"longitude"`
	City          *string  `json:"city"`
	Country       *string  `json:"country"`
	BirthPlace    *string  `json:"birth_place"`
	IsUnknownTime bool     `json:"is_unknown_time"`
	Note          *string  `json:"note"`
}

type UpdateBirthProfileRequest struct {
	FullName      *string  `json:"full_name" validate:"omitempty,min=1,max=150"`
	Gender        *string  `json:"gender"`
	BirthDate     *string  `json:"birth_date"`
	BirthTime     *string  `json:"birth_time"`
	Timezone      *string  `json:"timezone"`
	Latitude      *float64 `json:"latitude"`
	Longitude     *float64 `json:"longitude"`
	City          *string  `json:"city"`
	Country       *string  `json:"country"`
	BirthPlace    *string  `json:"birth_place"`
	IsUnknownTime *bool    `json:"is_unknown_time"`
	Note          *string  `json:"note"`
}

type BirthProfileResponse struct {
	ID            string   `json:"id"`
	UserID        string   `json:"user_id"`
	FullName      string   `json:"full_name"`
	Gender        *string  `json:"gender,omitempty"`
	BirthDate     string   `json:"birth_date"`
	BirthTime     *string  `json:"birth_time,omitempty"`
	Timezone      string   `json:"timezone"`
	Latitude      *float64 `json:"latitude,omitempty"`
	Longitude     *float64 `json:"longitude,omitempty"`
	City          *string  `json:"city,omitempty"`
	Country       *string  `json:"country,omitempty"`
	BirthPlace    *string  `json:"birth_place,omitempty"`
	IsUnknownTime bool     `json:"is_unknown_time"`
	Note          *string  `json:"note,omitempty"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}
