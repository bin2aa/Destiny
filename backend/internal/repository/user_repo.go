package repository

import (
	"context"
	"fmt"
	"time"

	"destiny-backend/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{pool: pool}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (name, email, password_hash, avatar, language, timezone, country)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at`

	return r.pool.QueryRow(ctx, query,
		user.Name, user.Email, user.PasswordHash,
		user.Avatar, user.Language, user.Timezone, user.Country,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (*model.User, error) {
	query := `SELECT id, name, email, password_hash, avatar, language, timezone, country, premium_until, created_at, updated_at
		FROM users WHERE id = $1`

	user := &model.User{}
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.PasswordHash,
		&user.Avatar, &user.Language, &user.Timezone, &user.Country,
		&user.PremiumUntil, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get user by id: %w", err)
	}
	return user, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, name, email, password_hash, avatar, language, timezone, country, premium_until, created_at, updated_at
		FROM users WHERE email = $1`

	user := &model.User{}
	err := r.pool.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.PasswordHash,
		&user.Avatar, &user.Language, &user.Timezone, &user.Country,
		&user.PremiumUntil, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *model.User) error {
	query := `UPDATE users SET name=$1, avatar=$2, language=$3, timezone=$4, country=$5, updated_at=now()
		WHERE id=$6 RETURNING updated_at`

	return r.pool.QueryRow(ctx, query,
		user.Name, user.Avatar, user.Language, user.Timezone, user.Country, user.ID,
	).Scan(&user.UpdatedAt)
}

func (r *UserRepo) UpdatePassword(ctx context.Context, userID, passwordHash string) error {
	query := `UPDATE users SET password_hash=$1, updated_at=now() WHERE id=$2`
	_, err := r.pool.Exec(ctx, query, passwordHash, userID)
	return err
}

func (r *UserRepo) UpdatePremiumUntil(ctx context.Context, userID string, until time.Time) error {
	query := `UPDATE users SET premium_until=$1, updated_at=now() WHERE id=$2`
	_, err := r.pool.Exec(ctx, query, until, userID)
	return err
}
