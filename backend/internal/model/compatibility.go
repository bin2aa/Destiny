package model

import "time"

type Compatibility struct {
	ID            string    `json:"id" db:"id"`
	ProfileAID    string    `json:"profile_a_id" db:"profile_a_id"`
	ProfileBID    string    `json:"profile_b_id" db:"profile_b_id"`
	LoveScore     *int16    `json:"love_score,omitempty" db:"love_score"`
	FriendScore   *int16    `json:"friend_score,omitempty" db:"friend_score"`
	MarriageScore *int16    `json:"marriage_score,omitempty" db:"marriage_score"`
	BusinessScore *int16    `json:"business_score,omitempty" db:"business_score"`
	Summary       *string   `json:"summary,omitempty" db:"summary"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
