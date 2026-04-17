package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/izukanji/ovc/pkg/utils"
)

func RequireRole(roles ...string) gin.HandlerFunc {
	allowed := make(map[string]bool, len(roles))
	for _, r := range roles {
		allowed[r] = true
	}
	return func(c *gin.Context) {
		role, _ := c.Get(UserRoleKey)
		if !allowed[role.(string)] {
			utils.Forbidden(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
