package model

import "time"

type PlanetPosition struct {
	ID             string    `json:"id" db:"id"`
	BirthProfileID string    `json:"birth_profile_id" db:"birth_profile_id"`
	PlanetID       int16     `json:"planet_id" db:"planet_id"`
	SignID         int16     `json:"sign_id" db:"sign_id"`
	House          *int16    `json:"house,omitempty" db:"house"`
	Degree         int16     `json:"degree" db:"degree"`
	Minute         int16     `json:"minute" db:"minute"`
	Retrograde     bool      `json:"retrograde" db:"retrograde"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type NatalAspect struct {
	ID             string    `json:"id" db:"id"`
	BirthProfileID string    `json:"birth_profile_id" db:"birth_profile_id"`
	PlanetAID      int16     `json:"planet_a_id" db:"planet_a_id"`
	PlanetBID      int16     `json:"planet_b_id" db:"planet_b_id"`
	AspectID       int16     `json:"aspect_id" db:"aspect_id"`
	Orb            float64   `json:"orb" db:"orb"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Lookup models
type ZodiacSign struct {
	ID          int16  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Element     string `json:"element" db:"element"`
	Modality    string `json:"modality" db:"modality"`
	Symbol      string `json:"symbol" db:"symbol"`
	Description string `json:"description" db:"description"`
}

type Planet struct {
	ID          int16  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Symbol      string `json:"symbol" db:"symbol"`
	Description string `json:"description" db:"description"`
}

type House struct {
	ID          int16  `json:"id" db:"id"`
	Number      int16  `json:"number" db:"number"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type Aspect struct {
	ID    int16   `json:"id" db:"id"`
	Name  string  `json:"name" db:"name"`
	Angle float64 `json:"angle" db:"angle"`
	Orb   float64 `json:"orb" db:"orb"`
}

type ChineseZodiac struct {
	ID          int16  `json:"id" db:"id"`
	Animal      string `json:"animal" db:"animal"`
	YinYang     string `json:"yin_yang" db:"yin_yang"`
	Element     string `json:"element" db:"element"`
	Description string `json:"description" db:"description"`
}

type FiveElement struct {
	ID          int16  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
