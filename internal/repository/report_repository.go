package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

const reportSelect = `
	SELECT id, title, body, term, year, created_by, created_at, updated_at
	FROM reports`

func scanReport(row *sql.Row) (*models.Report, error) {
	r := &models.Report{}
	err := row.Scan(&r.ID, &r.Title, &r.Body, &r.Term, &r.Year, &r.CreatedBy, &r.CreatedAt, &r.UpdatedAt)
	return r, err
}

func (r *ReportRepository) Create(ctx context.Context, report *models.Report) error {
	query := `
		INSERT INTO reports (title, body, term, year, created_by)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		report.Title, report.Body, report.Term, report.Year, report.CreatedBy,
	).Scan(&report.ID, &report.CreatedAt, &report.UpdatedAt)
}

func (r *ReportRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Report, error) {
	row := r.db.QueryRowContext(ctx, reportSelect+` WHERE id = $1`, id)
	report, err := scanReport(row)
	if err != nil {
		return nil, fmt.Errorf("report not found")
	}
	return report, nil
}

func (r *ReportRepository) List(ctx context.Context, filters models.ReportFilters) ([]models.Report, error) {
	query := reportSelect
	args := []interface{}{}
	i := 1

	if filters.Term != "" {
		query += fmt.Sprintf(" WHERE term = $%d", i)
		args = append(args, filters.Term)
		i++
		if filters.Year > 0 {
			query += fmt.Sprintf(" AND year = $%d", i)
			args = append(args, filters.Year)
			i++
		}
	} else if filters.Year > 0 {
		query += fmt.Sprintf(" WHERE year = $%d", i)
		args = append(args, filters.Year)
		i++
	}

	query += " ORDER BY created_at DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reports := []models.Report{}
	for rows.Next() {
		var rep models.Report
		if err := rows.Scan(&rep.ID, &rep.Title, &rep.Body, &rep.Term, &rep.Year, &rep.CreatedBy, &rep.CreatedAt, &rep.UpdatedAt); err != nil {
			return nil, err
		}
		reports = append(reports, rep)
	}
	return reports, rows.Err()
}

func (r *ReportRepository) Update(ctx context.Context, id uuid.UUID, req models.UpdateReportRequest) error {
	query := `
		UPDATE reports SET title = $1, body = $2, term = $3, year = $4, updated_at = NOW()
		WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, req.Title, req.Body, req.Term, req.Year, id)
	return err
}

func (r *ReportRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM reports WHERE id = $1`, id)
	return err
}
