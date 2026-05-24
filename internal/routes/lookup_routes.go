package routes

import (
	"net/http"

	"github.com/izukanji/ovc/internal/handlers"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/repository"
)

func registerLookupRoutes(mux *http.ServeMux, h *handlers.LookupHandler, authMW func(http.Handler) http.Handler, permRepo *repository.PermissionRepository) {
	lookupMW := middleware.RequirePermission(permRepo, "lookups", "read")

	mux.Handle("GET /api/categories", authMW(lookupMW(http.HandlerFunc(h.Categories))))
	mux.Handle("GET /api/requisites", authMW(lookupMW(http.HandlerFunc(h.Requisites))))
	mux.Handle("GET /api/sponsors", authMW(lookupMW(http.HandlerFunc(h.Sponsors))))
}
