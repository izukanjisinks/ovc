package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izukanji/ovc/internal/handlers"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/repository"
)

func registerUserRoutes(rg *gin.RouterGroup, h *handlers.UserHandler, authMW gin.HandlerFunc, permRepo *repository.PermissionRepository) {
	users := rg.Group("/users", authMW)
	users.POST("", middleware.RequirePermission(permRepo, "user-management", "create"), h.Create)
	users.GET("", middleware.RequirePermission(permRepo, "user-management", "read"), h.List)
	users.PUT("/:id", middleware.RequirePermission(permRepo, "user-management", "update"), h.Update)
	users.DELETE("/:id", middleware.RequirePermission(permRepo, "user-management", "delete"), h.Delete)
}
