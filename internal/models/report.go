package models

import (
	"time"

	"github.com/google/uuid"
)

type Term string

const (
	Term1 Term = "TERM_1"
	Term2 Term = "TERM_2"
	Term3 Term = "TERM_3"
)

type Report struct {
	ID            uuid.UUID      `json:"id"`
	Title         string         `json:"title"`
	Body          string         `json:"body"`
	Term          Term           `json:"term"`
	Year          int            `json:"year"`
	CreatedBy     *uuid.UUID     `json:"created_by"`
	CreatedByName string         `json:"created_by_name,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Beneficiaries []Child        `json:"beneficiaries,omitempty"`
}

type CreateReportRequest struct {
	Title    string      `json:"title"`
	Body     string      `json:"body"`
	Term     Term        `json:"term"`
	Year     int         `json:"year"`
	ChildIDs []uuid.UUID `json:"child_ids"`
}

type UpdateReportRequest struct {
	Title    string      `json:"title"`
	Body     string      `json:"body"`
	Term     Term        `json:"term"`
	Year     int         `json:"year"`
	ChildIDs []uuid.UUID `json:"child_ids"`
}

type ReportFilters struct {
	Term Term
	Year int
}

func (r CreateReportRequest) Validate() string {
	switch {
	case r.Title == "":
		return "title is required"
	case r.Body == "":
		return "body is required"
	case r.Term != Term1 && r.Term != Term2 && r.Term != Term3:
		return "term must be TERM_1, TERM_2 or TERM_3"
	case r.Year < 2000:
		return "valid year is required"
	}
	return ""
}

func (r UpdateReportRequest) Validate() string {
	switch {
	case r.Title == "":
		return "title is required"
	case r.Body == "":
		return "body is required"
	case r.Term != Term1 && r.Term != Term2 && r.Term != Term3:
		return "term must be TERM_1, TERM_2 or TERM_3"
	case r.Year < 2000:
		return "valid year is required"
	}
	return ""
}
