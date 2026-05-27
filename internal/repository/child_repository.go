package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/models"
)

type ChildRepository struct {
	db *sql.DB
}

func NewChildRepository(db *sql.DB) *ChildRepository {
	return &ChildRepository{db: db}
}

const childSelect = `
	SELECT id, pupil_id, first_name, last_name, address, class_name,
	       image_url, guardian_first_name, guardian_last_name,
	       guardian_address, guardian_phone, created_by, created_at, updated_at
	FROM children`

func scanChild(row *sql.Row) (*models.Child, error) {
	c := &models.Child{}
	err := row.Scan(
		&c.ID, &c.PupilID, &c.FirstName, &c.LastName, &c.Address, &c.ClassName,
		&c.ImageURL, &c.GuardianFirstName, &c.GuardianLastName,
		&c.GuardianAddress, &c.GuardianPhone, &c.CreatedBy, &c.CreatedAt, &c.UpdatedAt,
	)
	return c, err
}

func (r *ChildRepository) Create(ctx context.Context, c *models.Child) error {
	query := `
		INSERT INTO children
			(pupil_id, first_name, last_name, address, class_name, image_url,
			 guardian_first_name, guardian_last_name, guardian_address, guardian_phone, created_by)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		c.PupilID, c.FirstName, c.LastName, c.Address, c.ClassName, c.ImageURL,
		c.GuardianFirstName, c.GuardianLastName, c.GuardianAddress, c.GuardianPhone, c.CreatedBy,
	).Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt)
}

func (r *ChildRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Child, error) {
	row := r.db.QueryRowContext(ctx, childSelect+` WHERE id = $1`, id)
	c, err := scanChild(row)
	if err != nil {
		return nil, fmt.Errorf("child not found")
	}
	return c, nil
}

func (r *ChildRepository) List(ctx context.Context, search string) ([]models.Child, error) {
	query := childSelect
	args := []any{}
	if search != "" {
		query += ` WHERE pupil_id ILIKE $1`
		args = append(args, "%"+search+"%")
	}
	query += ` ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, args...)
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

func (r *ChildRepository) Update(ctx context.Context, id uuid.UUID, req models.UpdateChildRequest) error {
	query := `
		UPDATE children SET
			first_name = $1, last_name = $2, address = $3, class_name = $4,
			image_url = $5, guardian_first_name = $6, guardian_last_name = $7,
			guardian_address = $8, guardian_phone = $9, updated_at = NOW()
		WHERE id = $10`
	_, err := r.db.ExecContext(ctx, query,
		req.FirstName, req.LastName, req.Address, req.ClassName, req.ImageURL,
		req.GuardianFirstName, req.GuardianLastName, req.GuardianAddress, req.GuardianPhone, id,
	)
	return err
}

func (r *ChildRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM children WHERE id = $1`, id)
	return err
}

// --- Category relationships ---

func (r *ChildRepository) SetCategories(ctx context.Context, childID uuid.UUID, categoryIDs []uuid.UUID) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `DELETE FROM child_categories WHERE child_id = $1`, childID); err != nil {
		return err
	}
	for _, catID := range categoryIDs {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO child_categories (child_id, category_id) VALUES ($1, $2)`, childID, catID,
		); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *ChildRepository) GetCategories(ctx context.Context, childID uuid.UUID) ([]models.OVCCategory, error) {
	query := `
		SELECT c.id, c.name FROM ovc_categories c
		JOIN child_categories cc ON cc.category_id = c.id
		WHERE cc.child_id = $1 ORDER BY c.name`
	rows, err := r.db.QueryContext(ctx, query, childID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cats := []models.OVCCategory{}
	for rows.Next() {
		var cat models.OVCCategory
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	return cats, rows.Err()
}

// --- Requisite relationships ---

func (r *ChildRepository) SetRequisites(ctx context.Context, childID uuid.UUID, items []models.RequisiteItem) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `DELETE FROM child_requisites WHERE child_id = $1`, childID); err != nil {
		return err
	}
	for _, item := range items {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO child_requisites (child_id, requisite_id, quantity, checked, price_per_item)
			 VALUES ($1, $2, $3, $4, $5)`,
			childID, item.RequisiteID, item.Quantity, item.Checked, item.PricePerItem,
		); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *ChildRepository) GetRequisites(ctx context.Context, childID uuid.UUID) ([]models.ChildRequisite, error) {
	query := `
		SELECT r.id, r.name, cr.quantity, cr.checked, cr.price_per_item
		FROM requisites r
		JOIN child_requisites cr ON cr.requisite_id = r.id
		WHERE cr.child_id = $1 ORDER BY r.name`
	rows, err := r.db.QueryContext(ctx, query, childID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []models.ChildRequisite{}
	for rows.Next() {
		var item models.ChildRequisite
		if err := rows.Scan(&item.ID, &item.Name, &item.Quantity, &item.Checked, &item.PricePerItem); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

// --- Sponsor relationships ---

func (r *ChildRepository) SetSponsors(ctx context.Context, childID uuid.UUID, sponsorIDs []uuid.UUID) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `DELETE FROM child_sponsors WHERE child_id = $1`, childID); err != nil {
		return err
	}
	for _, sID := range sponsorIDs {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO child_sponsors (child_id, sponsor_id) VALUES ($1, $2)`, childID, sID,
		); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *ChildRepository) GetSponsors(ctx context.Context, childID uuid.UUID) ([]models.Sponsor, error) {
	query := `
		SELECT s.id, s.name FROM sponsors s
		JOIN child_sponsors cs ON cs.sponsor_id = s.id
		WHERE cs.child_id = $1 ORDER BY s.name`
	rows, err := r.db.QueryContext(ctx, query, childID)
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
