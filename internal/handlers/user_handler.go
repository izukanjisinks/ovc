package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/models"
	"github.com/izukanji/ovc/internal/services"
	"github.com/izukanji/ovc/pkg/utils"
)

type UserHandler struct {
	authService *services.AuthService
	userService *services.UserService
}

func NewUserHandler(authService *services.AuthService, userService *services.UserService) *UserHandler {
	return &UserHandler{authService: authService, userService: userService}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if req.Email == "" || req.Password == "" || req.FullName == "" {
		utils.BadRequest(w, "email, password and full_name are required")
		return
	}
	user, err := h.authService.Register(r.Context(), req)
	if err != nil {
		utils.BadRequest(w, err.Error())
		return
	}
	utils.Created(w, user)
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.List(r.Context())
	if err != nil {
		utils.InternalError(w)
		return
	}
	utils.OK(w, users)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid user id")
		return
	}
	var req models.UpdateUserRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.BadRequest(w, "invalid request body")
		return
	}
	if req.FullName == "" || req.RoleName == "" {
		utils.BadRequest(w, "full_name and role are required")
		return
	}
	if err := h.userService.Update(r.Context(), id, req); err != nil {
		utils.InternalError(w)
		return
	}
	utils.Message(w, "user updated")
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.BadRequest(w, "invalid user id")
		return
	}
	if err := h.userService.Delete(r.Context(), id); err != nil {
		utils.InternalError(w)
		return
	}
	utils.Message(w, "user deleted")
}
