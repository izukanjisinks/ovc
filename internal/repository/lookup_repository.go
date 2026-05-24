package repository

import (
	"context"
	"database/sql"

	"github.com/izukanji/ovc/internal/models"
)

type LookupRepository struct {
	db *sql.DB
}

func NewLookupRepository(db *sql.DB) *LookupRepository {
	return &LookupRepository{db: db}
}

func (r *LookupRepository) ListCategories(ctx context.Context) ([]models.OVCCategory, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name FROM ovc_categories ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []models.OVCCategory
	for rows.Next() {
		var c models.OVCCategory
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		cats = append(cats, c)
	}
	return cats, rows.Err()
}

func (r *LookupRepository) ListRequisites(ctx context.Context) ([]models.Requisite, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name, default_price FROM requisites ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Requisite
	for rows.Next() {
		var item models.Requisite
		if err := rows.Scan(&item.ID, &item.Name, &item.DefaultPrice); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *LookupRepository) ListSponsors(ctx context.Context) ([]models.Sponsor, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name FROM sponsors ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sponsors []models.Sponsor
	for rows.Next() {
		var s models.Sponsor
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		sponsors = append(sponsors, s)
	}
	return sponsors, rows.Err()
}
