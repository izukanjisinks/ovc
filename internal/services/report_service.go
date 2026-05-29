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
	if len(req.ChildIDs) > 0 {
		if err := s.reportRepo.SetBeneficiaries(ctx, report.ID, req.ChildIDs); err != nil {
			return nil, fmt.Errorf("failed to set beneficiaries")
		}
	}
	report.Beneficiaries, _ = s.reportRepo.GetBeneficiaries(ctx, report.ID)
	return report, nil
}

func (s *ReportService) GetByID(ctx context.Context, id uuid.UUID) (*models.Report, error) {
	report, err := s.reportRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	report.Beneficiaries, _ = s.reportRepo.GetBeneficiaries(ctx, id)
	return report, nil
}

func (s *ReportService) List(ctx context.Context, filters models.ReportFilters) ([]models.Report, error) {
	reports, err := s.reportRepo.List(ctx, filters)
	if err != nil {
		return nil, err
	}
	for i := range reports {
		reports[i].Beneficiaries, _ = s.reportRepo.GetBeneficiaries(ctx, reports[i].ID)
	}
	return reports, nil
}

func (s *ReportService) Update(ctx context.Context, id uuid.UUID, req models.UpdateReportRequest) error {
	if _, err := s.reportRepo.FindByID(ctx, id); err != nil {
		return fmt.Errorf("report not found")
	}
	if err := s.reportRepo.Update(ctx, id, req); err != nil {
		return err
	}
	return s.reportRepo.SetBeneficiaries(ctx, id, req.ChildIDs)
}

func (s *ReportService) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := s.reportRepo.FindByID(ctx, id); err != nil {
		return fmt.Errorf("report not found")
	}
	return s.reportRepo.Delete(ctx, id)
}
