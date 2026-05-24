package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/repository"
)

type ChildService struct {
	childRepo *repository.ChildRepository
}

func NewChildService(childRepo *repository.ChildRepository) *ChildService {
	return &ChildService{childRepo: childRepo}
}

func (s *ChildService) Create(ctx context.Context, req models.CreateChildRequest, createdBy uuid.UUID) (*models.Child, error) {
	child := &models.Child{
		PupilID:           req.PupilID,
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		Address:           req.Address,
		ClassName:         req.ClassName,
		ImageURL:          req.ImageURL,
		GuardianFirstName: req.GuardianFirstName,
		GuardianLastName:  req.GuardianLastName,
		GuardianAddress:   req.GuardianAddress,
		GuardianPhone:     req.GuardianPhone,
		CreatedBy:         &createdBy,
	}
	if err := s.childRepo.Create(ctx, child); err != nil {
		return nil, fmt.Errorf("pupil ID already exists")
	}
	return child, nil
}

func (s *ChildService) GetByID(ctx context.Context, id uuid.UUID) (*models.Child, error) {
	child, err := s.childRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	// Load relationships
	child.Categories, _ = s.childRepo.GetCategories(ctx, id)
	child.Requisites, _ = s.childRepo.GetRequisites(ctx, id)
	child.Sponsors, _ = s.childRepo.GetSponsors(ctx, id)
	return child, nil
}

func (s *ChildService) List(ctx context.Context, search string) ([]models.Child, error) {
	return s.childRepo.List(ctx, search)
}

func (s *ChildService) Update(ctx context.Context, id uuid.UUID, req models.UpdateChildRequest) error {
	if _, err := s.childRepo.FindByID(ctx, id); err != nil {
		return fmt.Errorf("child not found")
	}
	return s.childRepo.Update(ctx, id, req)
}

func (s *ChildService) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := s.childRepo.FindByID(ctx, id); err != nil {
		return fmt.Errorf("child not found")
	}
	return s.childRepo.Delete(ctx, id)
}

func (s *ChildService) SetCategories(ctx context.Context, childID uuid.UUID, req models.SetCategoriesRequest) error {
	return s.childRepo.SetCategories(ctx, childID, req.CategoryIDs)
}

func (s *ChildService) SetRequisites(ctx context.Context, childID uuid.UUID, req models.SetRequisitesRequest) error {
	return s.childRepo.SetRequisites(ctx, childID, req.Requisites)
}

func (s *ChildService) SetSponsors(ctx context.Context, childID uuid.UUID, req models.SetSponsorsRequest) error {
	return s.childRepo.SetSponsors(ctx, childID, req.SponsorIDs)
}
