package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
	permRepo *repository.PermissionRepository
}

func NewUserService(userRepo *repository.UserRepository, permRepo *repository.PermissionRepository) *UserService {
	return &UserService{userRepo: userRepo, permRepo: permRepo}
}

func (s *UserService) List(ctx context.Context) ([]models.User, error) {
	return s.userRepo.List(ctx)
}

func (s *UserService) Update(ctx context.Context, id uuid.UUID, req models.UpdateUserRequest) error {
	roleID, err := s.permRepo.RoleIDByName(ctx, req.RoleName)
	if err != nil {
		return err
	}
	return s.userRepo.Update(ctx, id, req.FullName, roleID)
}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.userRepo.Delete(ctx, id)
}
