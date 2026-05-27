package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/models"
)

type HighlightRepository struct {
	db *sql.DB
}

func NewHighlightRepository(db *sql.DB) *HighlightRepository {
	return &HighlightRepository{db: db}
}

func (r *HighlightRepository) Create(ctx context.Context, h *models.Highlight) error {
	query := `
		INSERT INTO highlights (image_url, caption, term, year)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`
	return r.db.QueryRowContext(ctx, query, h.ImageURL, h.Caption, h.Term, h.Year).
		Scan(&h.ID, &h.CreatedAt)
}

func (r *HighlightRepository) List(ctx context.Context) ([]models.Highlight, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, image_url, caption, term, year, created_at FROM highlights ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	highlights := []models.Highlight{}
	for rows.Next() {
		var h models.Highlight
		if err := rows.Scan(&h.ID, &h.ImageURL, &h.Caption, &h.Term, &h.Year, &h.CreatedAt); err != nil {
			return nil, err
		}
		highlights = append(highlights, h)
	}
	return highlights, rows.Err()
}

func (r *HighlightRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM highlights WHERE id = $1`, id)
	return err
}
