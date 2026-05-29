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

func (r *ReportRepository) SetBeneficiaries(ctx context.Context, reportID uuid.UUID, childIDs []uuid.UUID) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `DELETE FROM report_beneficiaries WHERE report_id = $1`, reportID); err != nil {
		return err
	}
	for _, childID := range childIDs {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO report_beneficiaries (report_id, child_id) VALUES ($1, $2)`, reportID, childID,
		); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *ReportRepository) GetBeneficiaries(ctx context.Context, reportID uuid.UUID) ([]models.Child, error) {
	query := `
		SELECT c.id, c.pupil_id, c.first_name, c.last_name, c.address, c.class_name,
		       c.image_url, c.guardian_first_name, c.guardian_last_name,
		       c.guardian_address, c.guardian_phone, c.created_by, c.created_at, c.updated_at
		FROM children c
		JOIN report_beneficiaries rb ON rb.child_id = c.id
		WHERE rb.report_id = $1
		ORDER BY c.last_name, c.first_name`

	rows, err := r.db.QueryContext(ctx, query, reportID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	children := []models.Child{}
	for rows.Next() {
		var c models.Child
		if err := rows.Scan(
			&c.ID, &c.PupilID, &c.FirstName, &c.LastName, &c.Address, &c.ClassName,
			&c.ImageURL, &c.GuardianFirstName, &c.GuardianLastName,
			&c.GuardianAddress, &c.GuardianPhone, &c.CreatedBy, &c.CreatedAt, &c.UpdatedAt,
		); err != nil {
			return nil, err
		}
		children = append(children, c)
	}
	return children, rows.Err()
}
