package repository

import (
	"context"
	"database/sql"

	"github.com/izukanji/ovc/internal/models"
)

type DashboardRepository struct {
	db *sql.DB
}

func NewDashboardRepository(db *sql.DB) *DashboardRepository {
	return &DashboardRepository{db: db}
}

func (r *DashboardRepository) Stats(ctx context.Context) (*models.DashboardStats, error) {
	stats := &models.DashboardStats{}

	if err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM children`).Scan(&stats.TotalChildren); err != nil {
		return nil, err
	}

	catRows, err := r.db.QueryContext(ctx, `
		SELECT oc.name, COUNT(cc.child_id)
		FROM ovc_categories oc
		LEFT JOIN child_categories cc ON cc.category_id = oc.id
		GROUP BY oc.name ORDER BY oc.name`)
	if err != nil {
		return nil, err
	}
	defer catRows.Close()
	for catRows.Next() {
		var s models.CategoryStat
		if err := catRows.Scan(&s.Category, &s.Count); err != nil {
			return nil, err
		}
		stats.ByCategory = append(stats.ByCategory, s)
	}

	spRows, err := r.db.QueryContext(ctx, `
		SELECT s.name, COUNT(cs.child_id)
		FROM sponsors s
		LEFT JOIN child_sponsors cs ON cs.sponsor_id = s.id
		GROUP BY s.name ORDER BY s.name`)
	if err != nil {
		return nil, err
	}
	defer spRows.Close()
	for spRows.Next() {
		var s models.SponsorStat
		if err := spRows.Scan(&s.Sponsor, &s.Count); err != nil {
			return nil, err
		}
		stats.BySponsor = append(stats.BySponsor, s)
	}

	return stats, nil
}
