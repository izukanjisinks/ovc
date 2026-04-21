package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/izukanji/ovc/internal/repository"
	"github.com/izukanji/ovc/pkg/utils"
)

func RequirePermission(permRepo *repository.PermissionRepository, resource, action string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			roleID, ok := r.Context().Value(UserRoleIDKey).(uuid.UUID)
			if !ok {
				utils.Forbidden(w)
				return
			}
			allowed, err := permRepo.HasPermission(r.Context(), roleID, resource, action)
			if err != nil || !allowed {
				utils.Forbidden(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
