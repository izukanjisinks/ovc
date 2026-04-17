package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/izukanji/ovc/internal/repository"
	"github.com/izukanji/ovc/pkg/utils"
)

func RequirePermission(permRepo *repository.PermissionRepository, resource, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, _ := c.Get(UserRoleIDKey)
		ok, err := permRepo.HasPermission(c.Request.Context(), roleID.(int), resource, action)
		if err != nil || !ok {
			utils.Forbidden(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
