package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izukanji/ovc/internal/handlers"
)

func registerAuthRoutes(rg *gin.RouterGroup, h *handlers.AuthHandler, authMW gin.HandlerFunc) {
	auth := rg.Group("/auth")
	auth.POST("/login", h.Login)
	auth.GET("/me", authMW, h.Me)
}
