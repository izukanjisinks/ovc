package handlers

import (
	"github.com/gin-gonic/gin"
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

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	resp, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		utils.Unauthorized(c, err.Error())
		return
	}
	utils.OK(c, resp)
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	user, err := h.authService.Me(c.Request.Context(), userID.(string))
	if err != nil {
		utils.NotFound(c, "user not found")
		return
	}
	utils.OK(c, user)
}
