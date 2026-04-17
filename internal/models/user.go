package models

import "time"

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	FullName     string    `json:"full_name"`
	RoleID       int       `json:"role_id"`
	RoleName     string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"full_name" binding:"required"`
	RoleName string `json:"role" binding:"required,oneof=admin user"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	RoleName string `json:"role" binding:"required,oneof=admin user"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
