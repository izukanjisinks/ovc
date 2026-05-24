package routes

import (
	"net/http"

	"github.com/izukanji/ovc/internal/handlers"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/repository"
)

func registerChildRoutes(mux *http.ServeMux, h *handlers.ChildHandler, authMW func(http.Handler) http.Handler, permRepo *repository.PermissionRepository) {
	mux.Handle("POST /api/children", authMW(
		middleware.RequirePermission(permRepo, "children", "create")(
			http.HandlerFunc(h.Create),
		),
	))
	mux.Handle("GET /api/children", authMW(
		middleware.RequirePermission(permRepo, "children", "read")(
			http.HandlerFunc(h.List),
		),
	))
	mux.Handle("GET /api/children/{id}", authMW(
		middleware.RequirePermission(permRepo, "children", "read")(
			http.HandlerFunc(h.Get),
		),
	))
	mux.Handle("PUT /api/children/{id}", authMW(
		middleware.RequirePermission(permRepo, "children", "update")(
			http.HandlerFunc(h.Update),
		),
	))
	mux.Handle("DELETE /api/children/{id}", authMW(
		middleware.RequirePermission(permRepo, "children", "delete")(
			http.HandlerFunc(h.Delete),
		),
	))

	// Relationship routes
	mux.Handle("PUT /api/children/{id}/categories", authMW(
		middleware.RequirePermission(permRepo, "children", "update")(
			http.HandlerFunc(h.SetCategories),
		),
	))
	mux.Handle("PUT /api/children/{id}/requisites", authMW(
		middleware.RequirePermission(permRepo, "children", "update")(
			http.HandlerFunc(h.SetRequisites),
		),
	))
	mux.Handle("PUT /api/children/{id}/sponsors", authMW(
		middleware.RequirePermission(permRepo, "children", "update")(
			http.HandlerFunc(h.SetSponsors),
		),
	))
}
