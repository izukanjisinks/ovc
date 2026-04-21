package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/services"
	"github.com/izukanji/ovc/pkg/utils"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if req.Email == "" || req.Password == "" {
		utils.BadRequest(w, "email and password are required")
		return
	}
	resp, err := h.authService.Login(r.Context(), req)
	if err != nil {
		utils.Unauthorized(w, err.Error())
		return
	}
	utils.OK(w, resp)
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		utils.Unauthorized(w, "unauthorized")
		return
	}
	user, err := h.authService.Me(r.Context(), userID)
	if err != nil {
		utils.NotFound(w, "user not found")
		return
	}
	utils.OK(w, user)
}
