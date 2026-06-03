package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/repository"
)

type LookupService struct {
	lookupRepo *repository.LookupRepository
}

func NewLookupService(lookupRepo *repository.LookupRepository) *LookupService {
	return &LookupService{lookupRepo: lookupRepo}
}

func (s *LookupService) Categories(ctx context.Context) ([]models.OVCCategory, error) {
	return s.lookupRepo.ListCategories(ctx)
}

func (s *LookupService) CreateCategory(ctx context.Context, req models.NameRequest) (*models.OVCCategory, error) {
	return s.lookupRepo.CreateCategory(ctx, req.Name)
}

func (s *LookupService) UpdateCategory(ctx context.Context, id uuid.UUID, req models.NameRequest) error {
	return s.lookupRepo.UpdateCategory(ctx, id, req.Name)
}

func (s *LookupService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	return s.lookupRepo.DeleteCategory(ctx, id)
}

func (s *LookupService) Requisites(ctx context.Context) ([]models.Requisite, error) {
	return s.lookupRepo.ListRequisites(ctx)
}

func (s *LookupService) CreateRequisite(ctx context.Context, req models.RequisiteRequest) (*models.Requisite, error) {
	return s.lookupRepo.CreateRequisite(ctx, req)
}

func (s *LookupService) UpdateRequisite(ctx context.Context, id uuid.UUID, req models.RequisiteRequest) error {
	return s.lookupRepo.UpdateRequisite(ctx, id, req)
}

func (s *LookupService) DeleteRequisite(ctx context.Context, id uuid.UUID) error {
	return s.lookupRepo.DeleteRequisite(ctx, id)
}

func (s *LookupService) Sponsors(ctx context.Context) ([]models.Sponsor, error) {
	return s.lookupRepo.ListSponsors(ctx)
}

func (s *LookupService) CreateSponsor(ctx context.Context, req models.NameRequest) (*models.Sponsor, error) {
	return s.lookupRepo.CreateSponsor(ctx, req.Name)
}

func (s *LookupService) UpdateSponsor(ctx context.Context, id uuid.UUID, req models.NameRequest) error {
	return s.lookupRepo.UpdateSponsor(ctx, id, req.Name)
}

func (s *LookupService) DeleteSponsor(ctx context.Context, id uuid.UUID) error {
	return s.lookupRepo.DeleteSponsor(ctx, id)
}
