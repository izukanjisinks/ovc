package handlers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/services"
	"github.com/izukanji/ovc/pkg/utils"
)

type ReportHandler struct {
	reportService *services.ReportService
}

func NewReportHandler(reportService *services.ReportService) *ReportHandler {
	return &ReportHandler{reportService: reportService}
}

func (h *ReportHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateReportRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if msg := req.Validate(); msg != "" {
		utils.BadRequest(w, msg)
		return
	}
	createdBy, _ := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	report, err := h.reportService.Create(r.Context(), req, createdBy)
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.Created(w, report)
}

func (h *ReportHandler) List(w http.ResponseWriter, r *http.Request) {
	filters := models.ReportFilters{
		Term: models.Term(r.URL.Query().Get("term")),
	}
	if y := r.URL.Query().Get("year"); y != "" {
		filters.Year, _ = strconv.Atoi(y)
	}
	reports, err := h.reportService.List(r.Context(), filters)
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.OK(w, reports)
}

func (h *ReportHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid report id")
		return
	}
	report, err := h.reportService.GetByID(r.Context(), id)
	if err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.OK(w, report)
}

func (h *ReportHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid report id")
		return
	}
	var req models.UpdateReportRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if msg := req.Validate(); msg != "" {
		utils.BadRequest(w, msg)
		return
	}
	if err := h.reportService.Update(r.Context(), id, req); err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.Message(w, "report updated")
}

func (h *ReportHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid report id")
		return
	}
	if err := h.reportService.Delete(r.Context(), id); err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.Message(w, "report deleted")
}
