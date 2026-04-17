package handlers

import (
	"github.com/gin-gonic/gin"
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

func (h *UserHandler) Create(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	user, err := h.authService.Register(c.Request.Context(), req)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	utils.Created(c, user)
}

func (h *UserHandler) List(c *gin.Context) {
	users, err := h.userService.List(c.Request.Context())
	if err != nil {
		utils.InternalError(c)
		return
	}
	utils.OK(c, users)
}

func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := h.userService.Update(c.Request.Context(), id, req); err != nil {
		utils.InternalError(c)
		return
	}
	utils.Message(c, "user updated")
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.userService.Delete(c.Request.Context(), id); err != nil {
		utils.InternalError(c)
		return
	}
	utils.Message(c, "user deleted")
}
