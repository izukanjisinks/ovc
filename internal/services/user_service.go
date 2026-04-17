package services

import (
	"context"

	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) List(ctx context.Context) ([]models.User, error) {
	return s.userRepo.List(ctx)
}

func (s *UserService) Update(ctx context.Context, id string, req models.UpdateUserRequest) error {
	return s.userRepo.Update(ctx, id, req.FullName, req.Role)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.userRepo.Delete(ctx, id)
}
