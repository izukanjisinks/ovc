package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/services"
	"github.com/izukanji/ovc/pkg/utils"
)

type ChildHandler struct {
	childService *services.ChildService
}

func NewChildHandler(childService *services.ChildService) *ChildHandler {
	return &ChildHandler{childService: childService}
}

func (h *ChildHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateChildRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if msg := req.Validate(); msg != "" {
		utils.BadRequest(w, msg)
		return
	}
	createdBy, _ := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	child, err := h.childService.Create(r.Context(), req, createdBy)
	if err != nil {
		utils.BadRequest(w, err.Error())
		return
	}
	utils.Created(w, child)
}

func (h *ChildHandler) List(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	children, err := h.childService.List(r.Context(), search)
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.OK(w, children)
}

func (h *ChildHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid child id")
		return
	}
	child, err := h.childService.GetByID(r.Context(), id)
	if err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.OK(w, child)
}

func (h *ChildHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid child id")
		return
	}
	var req models.UpdateChildRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if msg := req.Validate(); msg != "" {
		utils.BadRequest(w, msg)
		return
	}
	if err := h.childService.Update(r.Context(), id, req); err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.Message(w, "child updated")
}

func (h *ChildHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid child id")
		return
	}
	if err := h.childService.Delete(r.Context(), id); err != nil {
		utils.NotFound(w, err.Error())
		return
	}
	utils.Message(w, "child deleted")
}

func (h *ChildHandler) SetCategories(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid child id")
		return
	}
	var req models.SetCategoriesRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if err := h.childService.SetCategories(r.Context(), id, req); err != nil {
		utils.InternalError(w)
		return
	}
	utils.Message(w, "categories updated")
}

func (h *ChildHandler) SetRequisites(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid child id")
		return
	}
	var req models.SetRequisitesRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if err := h.childService.SetRequisites(r.Context(), id, req); err != nil {
		utils.InternalError(w)
		return
	}
	utils.Message(w, "requisites updated")
}

func (h *ChildHandler) SetSponsors(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid child id")
		return
	}
	var req models.SetSponsorsRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if err := h.childService.SetSponsors(r.Context(), id, req); err != nil {
		utils.InternalError(w)
		return
	}
	utils.Message(w, "sponsors updated")
}
