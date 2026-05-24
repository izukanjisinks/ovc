package services

import (
	"context"

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

func (s *LookupService) Requisites(ctx context.Context) ([]models.Requisite, error) {
	return s.lookupRepo.ListRequisites(ctx)
}

func (s *LookupService) Sponsors(ctx context.Context) ([]models.Sponsor, error) {
	return s.lookupRepo.ListSponsors(ctx)
}
