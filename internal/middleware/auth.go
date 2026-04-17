package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/izukanji/ovc/internal/config"
	"github.com/izukanji/ovc/pkg/utils"
)

const UserIDKey = "userID"
const UserRoleKey = "userRole"

func Auth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			utils.Unauthorized(c, "missing or invalid authorization header")
			c.Abort()
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		claims, err := utils.ValidateToken(token, cfg.JWTSecret)
		if err != nil {
			utils.Unauthorized(c, "invalid or expired token")
			c.Abort()
			return
		}
		c.Set(UserIDKey, claims.UserID)
		c.Set(UserRoleKey, string(claims.Role))
		c.Next()
	}
}
