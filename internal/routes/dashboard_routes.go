package routes

import (
	"net/http"

	"github.com/izukanji/ovc/internal/handlers"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/repository"
)

func registerDashboardRoutes(mux *http.ServeMux, dh *handlers.DashboardHandler, hh *handlers.HighlightHandler, authMW func(http.Handler) http.Handler, permRepo *repository.PermissionRepository) {
	mux.Handle("GET /api/v1/dashboard/stats", authMW(
		middleware.RequirePermission(permRepo, "dashboard", "read")(
			http.HandlerFunc(dh.Stats),
		),
	))
	mux.Handle("GET /api/v1/highlights", authMW(
		middleware.RequirePermission(permRepo, "highlights", "read")(
			http.HandlerFunc(hh.List),
		),
	))
	mux.Handle("POST /api/v1/highlights", authMW(
		middleware.RequirePermission(permRepo, "highlights", "create")(
			http.HandlerFunc(hh.Create),
		),
	))
	mux.Handle("DELETE /api/v1/highlights/{id}", authMW(
		middleware.RequirePermission(permRepo, "highlights", "delete")(
			http.HandlerFunc(hh.Delete),
		),
	))
}
