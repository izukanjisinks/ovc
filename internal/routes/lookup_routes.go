package routes

import (
	"net/http"

	"github.com/izukanji/ovc/internal/handlers"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/repository"
)

func registerLookupRoutes(mux *http.ServeMux, h *handlers.LookupHandler, authMW func(http.Handler) http.Handler, permRepo *repository.PermissionRepository) {
	readMW := middleware.RequirePermission(permRepo, "lookups", "read")
	createMW := middleware.RequirePermission(permRepo, "lookups", "create")
	updateMW := middleware.RequirePermission(permRepo, "lookups", "update")
	deleteMW := middleware.RequirePermission(permRepo, "lookups", "delete")

	mux.Handle("GET /api/v1/categories", authMW(readMW(http.HandlerFunc(h.Categories))))
	mux.Handle("POST /api/v1/categories", authMW(createMW(http.HandlerFunc(h.CreateCategory))))
	mux.Handle("PUT /api/v1/categories/{id}", authMW(updateMW(http.HandlerFunc(h.UpdateCategory))))
	mux.Handle("DELETE /api/v1/categories/{id}", authMW(deleteMW(http.HandlerFunc(h.DeleteCategory))))

	mux.Handle("GET /api/v1/requisites", authMW(readMW(http.HandlerFunc(h.Requisites))))
	mux.Handle("POST /api/v1/requisites", authMW(createMW(http.HandlerFunc(h.CreateRequisite))))
	mux.Handle("PUT /api/v1/requisites/{id}", authMW(updateMW(http.HandlerFunc(h.UpdateRequisite))))
	mux.Handle("DELETE /api/v1/requisites/{id}", authMW(deleteMW(http.HandlerFunc(h.DeleteRequisite))))

	mux.Handle("GET /api/v1/sponsors", authMW(readMW(http.HandlerFunc(h.Sponsors))))
	mux.Handle("POST /api/v1/sponsors", authMW(createMW(http.HandlerFunc(h.CreateSponsor))))
	mux.Handle("PUT /api/v1/sponsors/{id}", authMW(updateMW(http.HandlerFunc(h.UpdateSponsor))))
	mux.Handle("DELETE /api/v1/sponsors/{id}", authMW(deleteMW(http.HandlerFunc(h.DeleteSponsor))))
}
