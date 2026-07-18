package repository

import (
	"context"
	"fmt"

	"destiny-backend/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BirthProfileRepo struct {
	pool *pgxpool.Pool
}

func NewBirthProfileRepo(pool *pgxpool.Pool) *BirthProfileRepo {
	return &BirthProfileRepo{pool: pool}
}

func (r *BirthProfileRepo) Create(ctx context.Context, bp *model.BirthProfile) error {
	query := `INSERT INTO birth_profiles (user_id, full_name, gender, birth_date, birth_time, timezone,
		latitude, longitude, city, country, birth_place, is_unknown_time, note)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
		RETURNING id, created_at, updated_at`

	return r.pool.QueryRow(ctx, query,
		bp.UserID, bp.FullName, bp.Gender, bp.BirthDate, bp.BirthTime,
		bp.Timezone, bp.Latitude, bp.Longitude, bp.City, bp.Country,
		bp.BirthPlace, bp.IsUnknownTime, bp.Note,
	).Scan(&bp.ID, &bp.CreatedAt, &bp.UpdatedAt)
}

func (r *BirthProfileRepo) GetByID(ctx context.Context, id string) (*model.BirthProfile, error) {
	query := `SELECT id, user_id, full_name, gender, birth_date, birth_time, timezone,
		latitude, longitude, city, country, birth_place, is_unknown_time, note, created_at, updated_at
		FROM birth_profiles WHERE id = $1`

	bp := &model.BirthProfile{}
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&bp.ID, &bp.UserID, &bp.FullName, &bp.Gender, &bp.BirthDate, &bp.BirthTime,
		&bp.Timezone, &bp.Latitude, &bp.Longitude, &bp.City, &bp.Country,
		&bp.BirthPlace, &bp.IsUnknownTime, &bp.Note, &bp.CreatedAt, &bp.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get birth profile by id: %w", err)
	}
	return bp, nil
}

func (r *BirthProfileRepo) ListByUserID(ctx context.Context, userID string) ([]*model.BirthProfile, error) {
	query := `SELECT id, user_id, full_name, gender, birth_date, birth_time, timezone,
		latitude, longitude, city, country, birth_place, is_unknown_time, note, created_at, updated_at
		FROM birth_profiles WHERE user_id = $1 ORDER BY created_at DESC`

	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("list birth profiles: %w", err)
	}
	defer rows.Close()

	var profiles []*model.BirthProfile
	for rows.Next() {
		bp := &model.BirthProfile{}
		if err := rows.Scan(
			&bp.ID, &bp.UserID, &bp.FullName, &bp.Gender, &bp.BirthDate, &bp.BirthTime,
			&bp.Timezone, &bp.Latitude, &bp.Longitude, &bp.City, &bp.Country,
			&bp.BirthPlace, &bp.IsUnknownTime, &bp.Note, &bp.CreatedAt, &bp.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan birth profile: %w", err)
		}
		profiles = append(profiles, bp)
	}
	return profiles, nil
}

func (r *BirthProfileRepo) Update(ctx context.Context, bp *model.BirthProfile) error {
	query := `UPDATE birth_profiles SET full_name=$1, gender=$2, birth_date=$3, birth_time=$4,
		timezone=$5, latitude=$6, longitude=$7, city=$8, country=$9, birth_place=$10,
		is_unknown_time=$11, note=$12, updated_at=now()
		WHERE id=$13 AND user_id=$14 RETURNING updated_at`

	return r.pool.QueryRow(ctx, query,
		bp.FullName, bp.Gender, bp.BirthDate, bp.BirthTime,
		bp.Timezone, bp.Latitude, bp.Longitude, bp.City, bp.Country,
		bp.BirthPlace, bp.IsUnknownTime, bp.Note, bp.ID, bp.UserID,
	).Scan(&bp.UpdatedAt)
}

func (r *BirthProfileRepo) Delete(ctx context.Context, id, userID string) error {
	query := `DELETE FROM birth_profiles WHERE id = $1 AND user_id = $2`
	_, err := r.pool.Exec(ctx, query, id, userID)
	return err
}
