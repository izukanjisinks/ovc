package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izukanji/ovc/internal/config"
	"github.com/izukanji/ovc/internal/handlers"
	"github.com/izukanji/ovc/internal/middleware"
)

func registerUserRoutes(rg *gin.RouterGroup, h *handlers.UserHandler, cfg *config.Config) {
	users := rg.Group("/users", middleware.Auth(cfg), middleware.RequireRole("admin"))
	users.POST("", h.Create)
	users.GET("", h.List)
	users.PUT("/:id", h.Update)
	users.DELETE("/:id", h.Delete)
}
