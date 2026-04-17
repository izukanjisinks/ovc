package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/izukanji/ovc/internal/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
	query := `
		INSERT INTO users (email, password_hash, full_name, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at`
	return r.db.QueryRow(ctx, query, u.Email, u.PasswordHash, u.FullName, u.Role).
		Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	u := &models.User{}
	query := `SELECT id, email, password_hash, full_name, role, created_at, updated_at FROM users WHERE email = $1`
	err := r.db.QueryRow(ctx, query, email).
		Scan(&u.ID, &u.Email, &u.PasswordHash, &u.FullName, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	u := &models.User{}
	query := `SELECT id, email, password_hash, full_name, role, created_at, updated_at FROM users WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).
		Scan(&u.ID, &u.Email, &u.PasswordHash, &u.FullName, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (r *UserRepository) List(ctx context.Context) ([]models.User, error) {
	query := `SELECT id, email, full_name, role, created_at, updated_at FROM users ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Email, &u.FullName, &u.Role, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, id, fullName string, role models.Role) error {
	query := `UPDATE users SET full_name = $1, role = $2, updated_at = NOW() WHERE id = $3`
	_, err := r.db.Exec(ctx, query, fullName, role, id)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	return err
}
