package repository

import (
	"context"
	"fmt"

	"destiny-backend/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BillingRepo struct {
	pool *pgxpool.Pool
}

func NewBillingRepo(pool *pgxpool.Pool) *BillingRepo {
	return &BillingRepo{pool: pool}
}

func (r *BillingRepo) GetPlans(ctx context.Context) ([]*model.Plan, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, code, name, price, currency, duration_days, features, is_active, created_at, updated_at FROM plans WHERE is_active = true ORDER BY price")
	if err != nil {
		return nil, fmt.Errorf("get plans: %w", err)
	}
	defer rows.Close()

	var plans []*model.Plan
	for rows.Next() {
		p := &model.Plan{}
		if err := rows.Scan(&p.ID, &p.Code, &p.Name, &p.Price, &p.Currency, &p.DurationDays, &p.Features, &p.IsActive, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan plan: %w", err)
		}
		plans = append(plans, p)
	}
	return plans, nil
}

func (r *BillingRepo) GetSubscriptionsByUserID(ctx context.Context, userID string) ([]*model.Subscription, error) {
	query := `SELECT id, user_id, plan_id, start_at, end_at, status, created_at, updated_at
		FROM subscriptions WHERE user_id = $1 ORDER BY created_at DESC`

	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("get subscriptions: %w", err)
	}
	defer rows.Close()

	var subs []*model.Subscription
	for rows.Next() {
		s := &model.Subscription{}
		if err := rows.Scan(&s.ID, &s.UserID, &s.PlanID, &s.StartAt, &s.EndAt, &s.Status, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan subscription: %w", err)
		}
		subs = append(subs, s)
	}
	return subs, nil
}

func (r *BillingRepo) GetActiveSubscription(ctx context.Context, userID string) (*model.Subscription, error) {
	query := `SELECT id, user_id, plan_id, start_at, end_at, status, created_at, updated_at
		FROM subscriptions WHERE user_id = $1 AND status = 'active' AND end_at > now() ORDER BY end_at DESC LIMIT 1`

	s := &model.Subscription{}
	err := r.pool.QueryRow(ctx, query, userID).Scan(
		&s.ID, &s.UserID, &s.PlanID, &s.StartAt, &s.EndAt, &s.Status, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get active subscription: %w", err)
	}
	return s, nil
}

func (r *BillingRepo) CreateSubscription(ctx context.Context, sub *model.Subscription) error {
	query := `INSERT INTO subscriptions (user_id, plan_id, start_at, end_at, status)
		VALUES ($1,$2,$3,$4,$5) RETURNING id, created_at, updated_at`
	return r.pool.QueryRow(ctx, query, sub.UserID, sub.PlanID, sub.StartAt, sub.EndAt, sub.Status).
		Scan(&sub.ID, &sub.CreatedAt, &sub.UpdatedAt)
}

func (r *BillingRepo) CreatePayment(ctx context.Context, p *model.Payment) error {
	query := `INSERT INTO payments (subscription_id, provider, amount, currency, status, transaction_id)
		VALUES ($1,$2,$3,$4,$5,$6) RETURNING id, created_at, updated_at`
	return r.pool.QueryRow(ctx, query, p.SubscriptionID, p.Provider, p.Amount, p.Currency, p.Status, p.TransactionID).
		Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}
