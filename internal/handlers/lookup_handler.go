package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/services"
	"github.com/izukanji/ovc/pkg/utils"
)

type LookupHandler struct {
	lookupService *services.LookupService
}

func NewLookupHandler(lookupService *services.LookupService) *LookupHandler {
	return &LookupHandler{lookupService: lookupService}
}

// --- Categories ---

func (h *LookupHandler) Categories(w http.ResponseWriter, r *http.Request) {
	cats, err := h.lookupService.Categories(r.Context())
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.OK(w, cats)
}

func (h *LookupHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req models.NameRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if msg := req.Validate(); msg != "" {
		utils.BadRequest(w, msg)
		return
	}
	cat, err := h.lookupService.CreateCategory(r.Context(), req)
	if err != nil {
		utils.BadRequest(w, "name already exists")
		return
	}
	utils.Created(w, cat)
}

func (h *LookupHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid id")
		return
	}
	var req models.NameRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if msg := req.Validate(); msg != "" {
		utils.BadRequest(w, msg)
		return
	}
	if err := h.lookupService.UpdateCategory(r.Context(), id, req); err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.Message(w, "category updated")
}

func (h *LookupHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid id")
		return
	}
	if err := h.lookupService.DeleteCategory(r.Context(), id); err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.Message(w, "category deleted")
}

// --- Requisites ---

func (h *LookupHandler) Requisites(w http.ResponseWriter, r *http.Request) {
	items, err := h.lookupService.Requisites(r.Context())
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.OK(w, items)
}

func (h *LookupHandler) CreateRequisite(w http.ResponseWriter, r *http.Request) {
	var req models.RequisiteRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if msg := req.Validate(); msg != "" {
		utils.BadRequest(w, msg)
		return
	}
	item, err := h.lookupService.CreateRequisite(r.Context(), req)
	if err != nil {
		utils.BadRequest(w, "name already exists")
		return
	}
	utils.Created(w, item)
}

func (h *LookupHandler) UpdateRequisite(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid id")
		return
	}
	var req models.RequisiteRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if msg := req.Validate(); msg != "" {
		utils.BadRequest(w, msg)
		return
	}
	if err := h.lookupService.UpdateRequisite(r.Context(), id, req); err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.Message(w, "requisite updated")
}

func (h *LookupHandler) DeleteRequisite(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid id")
		return
	}
	if err := h.lookupService.DeleteRequisite(r.Context(), id); err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.Message(w, "requisite deleted")
}

// --- Sponsors ---

func (h *LookupHandler) Sponsors(w http.ResponseWriter, r *http.Request) {
	sponsors, err := h.lookupService.Sponsors(r.Context())
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.OK(w, sponsors)
}

func (h *LookupHandler) CreateSponsor(w http.ResponseWriter, r *http.Request) {
	var req models.NameRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if msg := req.Validate(); msg != "" {
		utils.BadRequest(w, msg)
		return
	}
	sponsor, err := h.lookupService.CreateSponsor(r.Context(), req)
	if err != nil {
		utils.BadRequest(w, "name already exists")
		return
	}
	utils.Created(w, sponsor)
}

func (h *LookupHandler) UpdateSponsor(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid id")
		return
	}
	var req models.NameRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if msg := req.Validate(); msg != "" {
		utils.BadRequest(w, msg)
		return
	}
	if err := h.lookupService.UpdateSponsor(r.Context(), id, req); err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.Message(w, "sponsor updated")
}

func (h *LookupHandler) DeleteSponsor(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid id")
		return
	}
	if err := h.lookupService.DeleteSponsor(r.Context(), id); err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.Message(w, "sponsor deleted")
}
