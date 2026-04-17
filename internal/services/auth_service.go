package services

import (
	"context"
	"fmt"

	"github.com/izukanji/ovc/internal/config"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/repository"
	"github.com/izukanji/ovc/pkg/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
	cfg      *config.Config
}

func NewAuthService(userRepo *repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{userRepo: userRepo, cfg: cfg}
}

func (s *AuthService) Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, fmt.Errorf("invalid email or password")
	}
	token, err := utils.GenerateToken(user.ID, string(user.Role), s.cfg.JWTSecret, s.cfg.JWTExpiry)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}
	return &models.LoginResponse{Token: token, User: *user}, nil
}

func (s *AuthService) Register(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password")
	}
	role := req.Role
	if role == "" {
		role = models.RoleUser
	}
	user := &models.User{
		Email:        req.Email,
		PasswordHash: hash,
		FullName:     req.FullName,
		Role:         role,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("email already in use")
	}
	return user, nil
}

func (s *AuthService) Me(ctx context.Context, userID string) (*models.User, error) {
	return s.userRepo.FindByID(ctx, userID)
}
