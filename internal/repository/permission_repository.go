package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type PermissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) HasPermission(ctx context.Context, roleID uuid.UUID, resource, action string) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1 FROM role_permissions
			WHERE role_id = $1 AND resource = $2 AND action = $3
		)`
	err := r.db.QueryRowContext(ctx, query, roleID, resource, action).Scan(&exists)
	return exists, err
}

func (r *PermissionRepository) RoleIDByName(ctx context.Context, name string) (uuid.UUID, error) {
	var id uuid.UUID
	err := r.db.QueryRowContext(ctx, `SELECT id FROM roles WHERE name = $1`, name).Scan(&id)
	return id, err
}
