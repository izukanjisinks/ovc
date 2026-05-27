package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/repository"
)

type HighlightService struct {
	highlightRepo *repository.HighlightRepository
}

func NewHighlightService(highlightRepo *repository.HighlightRepository) *HighlightService {
	return &HighlightService{highlightRepo: highlightRepo}
}

func (s *HighlightService) Create(ctx context.Context, req models.CreateHighlightRequest) (*models.Highlight, error) {
	h := &models.Highlight{
		ImageURL: req.ImageURL,
		Caption:  req.Caption,
		Term:     req.Term,
		Year:     req.Year,
	}
	if err := s.highlightRepo.Create(ctx, h); err != nil {
		return nil, err
	}
	return h, nil
}

func (s *HighlightService) List(ctx context.Context) ([]models.Highlight, error) {
	return s.highlightRepo.List(ctx)
}

func (s *HighlightService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.highlightRepo.Delete(ctx, id)
}
