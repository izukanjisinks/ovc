package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/repository"
)

type ReportService struct {
	reportRepo *repository.ReportRepository
}

func NewReportService(reportRepo *repository.ReportRepository) *ReportService {
	return &ReportService{reportRepo: reportRepo}
}

func (s *ReportService) Create(ctx context.Context, req models.CreateReportRequest, createdBy uuid.UUID) (*models.Report, error) {
	report := &models.Report{
		Title:     req.Title,
		Body:      req.Body,
		Term:      req.Term,
		Year:      req.Year,
		CreatedBy: &createdBy,
	}
	if err := s.reportRepo.Create(ctx, report); err != nil {
		return nil, fmt.Errorf("failed to create report")
	}
	return report, nil
}

func (s *ReportService) GetByID(ctx context.Context, id uuid.UUID) (*models.Report, error) {
	return s.reportRepo.FindByID(ctx, id)
}

func (s *ReportService) List(ctx context.Context, filters models.ReportFilters) ([]models.Report, error) {
	return s.reportRepo.List(ctx, filters)
}

func (s *ReportService) Update(ctx context.Context, id uuid.UUID, req models.UpdateReportRequest) error {
	if _, err := s.reportRepo.FindByID(ctx, id); err != nil {
		return fmt.Errorf("report not found")
	}
	return s.reportRepo.Update(ctx, id, req)
}

func (s *ReportService) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := s.reportRepo.FindByID(ctx, id); err != nil {
		return fmt.Errorf("report not found")
	}
	return s.reportRepo.Delete(ctx, id)
}
