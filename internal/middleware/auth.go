package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/config"
	"github.com/izukanji/ovc/internal/repository"
	"github.com/izukanji/ovc/pkg/utils"
)

const UserIDKey = "userID"
const UserRoleIDKey = "userRoleID"

func Auth(cfg *config.Config, userRepo *repository.UserRepository) gin.HandlerFunc {
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
		userID, err := uuid.Parse(claims.UserID)
		if err != nil {
			utils.Unauthorized(c, "invalid token subject")
			c.Abort()
			return
		}
		user, err := userRepo.FindByID(c.Request.Context(), userID)
		if err != nil {
			utils.Unauthorized(c, "user not found")
			c.Abort()
			return
		}
		c.Set(UserIDKey, user.ID)
		c.Set(UserRoleIDKey, user.RoleID)
		c.Next()
	}
}
