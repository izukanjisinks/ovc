package routes

import (
	"net/http"

	"github.com/izukanji/ovc/internal/handlers"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/repository"
)

func registerReportRoutes(mux *http.ServeMux, h *handlers.ReportHandler, authMW func(http.Handler) http.Handler, permRepo *repository.PermissionRepository) {
	mux.Handle("POST /api/v1/reports", authMW(
		middleware.RequirePermission(permRepo, "reports", "create")(
			http.HandlerFunc(h.Create),
		),
	))
	mux.Handle("GET /api/v1/reports", authMW(
		middleware.RequirePermission(permRepo, "reports", "read")(
			http.HandlerFunc(h.List),
		),
	))
	mux.Handle("GET /api/v1/reports/{id}", authMW(
		middleware.RequirePermission(permRepo, "reports", "read")(
			http.HandlerFunc(h.Get),
		),
	))
	mux.Handle("PUT /api/v1/reports/{id}", authMW(
		middleware.RequirePermission(permRepo, "reports", "update")(
			http.HandlerFunc(h.Update),
		),
	))
	mux.Handle("DELETE /api/v1/reports/{id}", authMW(
		middleware.RequirePermission(permRepo, "reports", "delete")(
			http.HandlerFunc(h.Delete),
		),
	))
}
