package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/izukanji/ovc/internal/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

const userSelect = `
	SELECT u.id, u.email, u.password_hash, u.full_name, u.role_id, r.name, u.created_at, u.updated_at
	FROM users u
	JOIN roles r ON r.id = u.role_id`

func scanUser(row interface{ Scan(...any) error }) (*models.User, error) {
	u := &models.User{}
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.FullName, &u.RoleID, &u.RoleName, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
	query := `
		INSERT INTO users (email, password_hash, full_name, role_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at`
	return r.db.QueryRow(ctx, query, u.Email, u.PasswordHash, u.RoleID).
		Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	row := r.db.QueryRow(ctx, userSelect+` WHERE u.email = $1`, email)
	u, err := scanUser(row)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	row := r.db.QueryRow(ctx, userSelect+` WHERE u.id = $1`, id)
	u, err := scanUser(row)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (r *UserRepository) List(ctx context.Context) ([]models.User, error) {
	query := `
		SELECT u.id, u.email, u.full_name, u.role_id, r.name, u.created_at, u.updated_at
		FROM users u
		JOIN roles r ON r.id = u.role_id
		ORDER BY u.created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Email, &u.FullName, &u.RoleID, &u.RoleName, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, id uuid.UUID, fullName string, roleID uuid.UUID) error {
	_, err := r.db.Exec(ctx,
		`UPDATE users SET full_name = $1, role_id = $2, updated_at = NOW() WHERE id = $3`,
		fullName, roleID, id,
	)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	return err
}
