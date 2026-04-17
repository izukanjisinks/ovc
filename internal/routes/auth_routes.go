package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izukanji/ovc/internal/config"
	"github.com/izukanji/ovc/internal/handlers"
	"github.com/izukanji/ovc/internal/middleware"
)

func registerAuthRoutes(rg *gin.RouterGroup, h *handlers.AuthHandler, cfg *config.Config) {
	auth := rg.Group("/auth")
	auth.POST("/login", h.Login)
	auth.GET("/me", middleware.Auth(cfg), h.Me)
}
