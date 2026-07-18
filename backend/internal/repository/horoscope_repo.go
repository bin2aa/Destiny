package repository

import (
	"context"
	"fmt"

	"destiny-backend/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HoroscopeRepo struct {
	pool *pgxpool.Pool
}

func NewHoroscopeRepo(pool *pgxpool.Pool) *HoroscopeRepo {
	return &HoroscopeRepo{pool: pool}
}

func (r *HoroscopeRepo) GetByProfileAndDate(ctx context.Context, profileID string, date string) (*model.DailyHoroscope, error) {
	query := `SELECT id, birth_profile_id, date, summary, career, love, health, finance, lucky_color, lucky_number, created_at
		FROM daily_horoscope WHERE birth_profile_id = $1 AND date = $2`

	h := &model.DailyHoroscope{}
	err := r.pool.QueryRow(ctx, query, profileID, date).Scan(
		&h.ID, &h.BirthProfileID, &h.Date, &h.Summary, &h.Career, &h.Love, &h.Health, &h.Finance, &h.LuckyColor, &h.LuckyNumber, &h.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get horoscope: %w", err)
	}
	return h, nil
}

func (r *HoroscopeRepo) Upsert(ctx context.Context, h *model.DailyHoroscope) error {
	query := `INSERT INTO daily_horoscope (birth_profile_id, date, summary, career, love, health, finance, lucky_color, lucky_number)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		ON CONFLICT (birth_profile_id, date) DO UPDATE SET
		summary=EXCLUDED.summary, career=EXCLUDED.career, love=EXCLUDED.love,
		health=EXCLUDED.health, finance=EXCLUDED.finance,
		lucky_color=EXCLUDED.lucky_color, lucky_number=EXCLUDED.lucky_number
		RETURNING id, created_at`
	return r.pool.QueryRow(ctx, query,
		h.BirthProfileID, h.Date, h.Summary, h.Career, h.Love, h.Health, h.Finance, h.LuckyColor, h.LuckyNumber,
	).Scan(&h.ID, &h.CreatedAt)
}
