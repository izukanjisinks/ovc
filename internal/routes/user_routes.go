package routes

import (
	"net/http"

	"github.com/izukanji/ovc/internal/handlers"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/repository"
)

func registerUserRoutes(mux *http.ServeMux, h *handlers.UserHandler, authMW func(http.Handler) http.Handler, permRepo *repository.PermissionRepository) {
	mux.Handle("POST /api/users", authMW(
		middleware.RequirePermission(permRepo, "user-management", "create")(
			http.HandlerFunc(h.Create),
		),
	))
	mux.Handle("GET /api/users", authMW(
		middleware.RequirePermission(permRepo, "user-management", "read")(
			http.HandlerFunc(h.List),
		),
	))
	mux.Handle("PUT /api/users/{id}", authMW(
		middleware.RequirePermission(permRepo, "user-management", "update")(
			http.HandlerFunc(h.Update),
		),
	))
	mux.Handle("DELETE /api/users/{id}", authMW(
		middleware.RequirePermission(permRepo, "user-management", "delete")(
			http.HandlerFunc(h.Delete),
		),
	))
}
