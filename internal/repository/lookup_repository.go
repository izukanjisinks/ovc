package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
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

	cats := []models.OVCCategory{}
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

	items := []models.Requisite{}
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

	sponsors := []models.Sponsor{}
	for rows.Next() {
		var s models.Sponsor
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		sponsors = append(sponsors, s)
	}
	return sponsors, rows.Err()
}

// --- Categories CRUD ---

func (r *LookupRepository) CreateCategory(ctx context.Context, name string) (*models.OVCCategory, error) {
	c := &models.OVCCategory{}
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO ovc_categories (name) VALUES ($1) RETURNING id, name`, name,
	).Scan(&c.ID, &c.Name)
	return c, err
}

func (r *LookupRepository) UpdateCategory(ctx context.Context, id uuid.UUID, name string) error {
	res, err := r.db.ExecContext(ctx, `UPDATE ovc_categories SET name = $1 WHERE id = $2`, name, id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("category not found")
	}
	return nil
}

func (r *LookupRepository) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM ovc_categories WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("category not found")
	}
	return nil
}

// --- Requisites CRUD ---

func (r *LookupRepository) CreateRequisite(ctx context.Context, req models.RequisiteRequest) (*models.Requisite, error) {
	item := &models.Requisite{}
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO requisites (name, default_price) VALUES ($1, $2) RETURNING id, name, default_price`,
		req.Name, req.DefaultPrice,
	).Scan(&item.ID, &item.Name, &item.DefaultPrice)
	return item, err
}

func (r *LookupRepository) UpdateRequisite(ctx context.Context, id uuid.UUID, req models.RequisiteRequest) error {
	res, err := r.db.ExecContext(ctx,
		`UPDATE requisites SET name = $1, default_price = $2 WHERE id = $3`,
		req.Name, req.DefaultPrice, id,
	)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("requisite not found")
	}
	return nil
}

func (r *LookupRepository) DeleteRequisite(ctx context.Context, id uuid.UUID) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM requisites WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("requisite not found")
	}
	return nil
}

// --- Sponsors CRUD ---

func (r *LookupRepository) CreateSponsor(ctx context.Context, name string) (*models.Sponsor, error) {
	s := &models.Sponsor{}
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO sponsors (name) VALUES ($1) RETURNING id, name`, name,
	).Scan(&s.ID, &s.Name)
	return s, err
}

func (r *LookupRepository) UpdateSponsor(ctx context.Context, id uuid.UUID, name string) error {
	res, err := r.db.ExecContext(ctx, `UPDATE sponsors SET name = $1 WHERE id = $2`, name, id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("sponsor not found")
	}
	return nil
}

func (r *LookupRepository) DeleteSponsor(ctx context.Context, id uuid.UUID) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM sponsors WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("sponsor not found")
	}
	return nil
}
