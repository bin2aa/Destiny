package repository

import (
	"context"
	"fmt"

	"destiny-backend/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LookupRepo struct {
	pool *pgxpool.Pool
}

func NewLookupRepo(pool *pgxpool.Pool) *LookupRepo {
	return &LookupRepo{pool: pool}
}

func (r *LookupRepo) GetZodiacSigns(ctx context.Context) ([]*model.ZodiacSign, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, name, element, modality, symbol, description FROM zodiac_signs ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("get zodiac signs: %w", err)
	}
	defer rows.Close()

	var result []*model.ZodiacSign
	for rows.Next() {
		z := &model.ZodiacSign{}
		if err := rows.Scan(&z.ID, &z.Name, &z.Element, &z.Modality, &z.Symbol, &z.Description); err != nil {
			return nil, fmt.Errorf("scan zodiac sign: %w", err)
		}
		result = append(result, z)
	}
	return result, nil
}

func (r *LookupRepo) GetPlanets(ctx context.Context) ([]*model.Planet, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, name, symbol, description FROM planets ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("get planets: %w", err)
	}
	defer rows.Close()

	var result []*model.Planet
	for rows.Next() {
		p := &model.Planet{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Symbol, &p.Description); err != nil {
			return nil, fmt.Errorf("scan planet: %w", err)
		}
		result = append(result, p)
	}
	return result, nil
}

func (r *LookupRepo) GetHouses(ctx context.Context) ([]*model.House, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, number, name, description FROM houses ORDER BY number")
	if err != nil {
		return nil, fmt.Errorf("get houses: %w", err)
	}
	defer rows.Close()

	var result []*model.House
	for rows.Next() {
		h := &model.House{}
		if err := rows.Scan(&h.ID, &h.Number, &h.Name, &h.Description); err != nil {
			return nil, fmt.Errorf("scan house: %w", err)
		}
		result = append(result, h)
	}
	return result, nil
}

func (r *LookupRepo) GetAspects(ctx context.Context) ([]*model.Aspect, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, name, angle, orb FROM aspects ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("get aspects: %w", err)
	}
	defer rows.Close()

	var result []*model.Aspect
	for rows.Next() {
		a := &model.Aspect{}
		if err := rows.Scan(&a.ID, &a.Name, &a.Angle, &a.Orb); err != nil {
			return nil, fmt.Errorf("scan aspect: %w", err)
		}
		result = append(result, a)
	}
	return result, nil
}

func (r *LookupRepo) GetChineseZodiac(ctx context.Context) ([]*model.ChineseZodiac, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, animal, yin_yang, element, description FROM chinese_zodiac ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("get chinese zodiac: %w", err)
	}
	defer rows.Close()

	var result []*model.ChineseZodiac
	for rows.Next() {
		c := &model.ChineseZodiac{}
		if err := rows.Scan(&c.ID, &c.Animal, &c.YinYang, &c.Element, &c.Description); err != nil {
			return nil, fmt.Errorf("scan chinese zodiac: %w", err)
		}
		result = append(result, c)
	}
	return result, nil
}

func (r *LookupRepo) GetFiveElements(ctx context.Context) ([]*model.FiveElement, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, name, description FROM five_elements ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("get five elements: %w", err)
	}
	defer rows.Close()

	var result []*model.FiveElement
	for rows.Next() {
		f := &model.FiveElement{}
		if err := rows.Scan(&f.ID, &f.Name, &f.Description); err != nil {
			return nil, fmt.Errorf("scan five element: %w", err)
		}
		result = append(result, f)
	}
	return result, nil
}
