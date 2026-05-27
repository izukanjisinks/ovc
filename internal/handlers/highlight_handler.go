package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/services"
	"github.com/izukanji/ovc/pkg/utils"
)

type HighlightHandler struct {
	highlightService *services.HighlightService
}

func NewHighlightHandler(highlightService *services.HighlightService) *HighlightHandler {
	return &HighlightHandler{highlightService: highlightService}
}

func (h *HighlightHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateHighlightRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if msg := req.Validate(); msg != "" {
		utils.BadRequest(w, msg)
		return
	}
	highlight, err := h.highlightService.Create(r.Context(), req)
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.Created(w, highlight)
}

func (h *HighlightHandler) List(w http.ResponseWriter, r *http.Request) {
	highlights, err := h.highlightService.List(r.Context())
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.OK(w, highlights)
}

func (h *HighlightHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid highlight id")
		return
	}
	if err := h.highlightService.Delete(r.Context(), id); err != nil {
		utils.InternalError(w)
		return
	}
	utils.Message(w, "highlight deleted")
}
