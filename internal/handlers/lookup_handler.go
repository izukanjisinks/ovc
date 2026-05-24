package handlers

import (
	"net/http"

	"github.com/izukanji/ovc/internal/services"
	"github.com/izukanji/ovc/pkg/utils"
)

type LookupHandler struct {
	lookupService *services.LookupService
}

func NewLookupHandler(lookupService *services.LookupService) *LookupHandler {
	return &LookupHandler{lookupService: lookupService}
}

func (h *LookupHandler) Categories(w http.ResponseWriter, r *http.Request) {
	cats, err := h.lookupService.Categories(r.Context())
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.OK(w, cats)
}

func (h *LookupHandler) Requisites(w http.ResponseWriter, r *http.Request) {
	items, err := h.lookupService.Requisites(r.Context())
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.OK(w, items)
}

func (h *LookupHandler) Sponsors(w http.ResponseWriter, r *http.Request) {
	sponsors, err := h.lookupService.Sponsors(r.Context())
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.OK(w, sponsors)
}
