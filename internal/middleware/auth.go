package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/config"
	"github.com/izukanji/ovc/internal/repository"
	"github.com/izukanji/ovc/pkg/utils"
)

type contextKey string

const (
	UserIDKey     contextKey = "userID"
	UserRoleIDKey contextKey = "userRoleID"
)

func Auth(cfg *config.Config, userRepo *repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if !strings.HasPrefix(header, "Bearer ") {
				utils.Unauthorized(w, "missing or invalid authorization header")
				return
			}
			token := strings.TrimPrefix(header, "Bearer ")
			claims, err := utils.ValidateToken(token, cfg.JWTSecret)
			if err != nil {
				utils.Unauthorized(w, "invalid or expired token")
				return
			}
			userID, err := uuid.Parse(claims.UserID)
			if err != nil {
				utils.Unauthorized(w, "invalid token subject")
				return
			}
			user, err := userRepo.FindByID(r.Context(), userID)
			if err != nil {
				utils.Unauthorized(w, "user not found")
				return
			}
			ctx := context.WithValue(r.Context(), UserIDKey, user.ID)
			ctx = context.WithValue(ctx, UserRoleIDKey, user.RoleID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
