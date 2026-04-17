package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PermissionRepository struct {
	db *pgxpool.Pool
}

func NewPermissionRepository(db *pgxpool.Pool) *PermissionRepository {
	return &PermissionRepository{db: db}
}

// HasPermission checks if a role has access to resource+action in one query — no JOIN needed.
func (r *PermissionRepository) HasPermission(ctx context.Context, roleID int, resource, action string) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1 FROM role_permissions
			WHERE role_id = $1 AND resource = $2 AND action = $3
		)`
	err := r.db.QueryRow(ctx, query, roleID, resource, action).Scan(&exists)
	return exists, err
}

// RoleIDByName returns the integer ID for a role name.
func (r *PermissionRepository) RoleIDByName(ctx context.Context, name string) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, `SELECT id FROM roles WHERE name = $1`, name).Scan(&id)
	return id, err
}
