package routes

import (
	"net/http"

	"github.com/izukanji/ovc/internal/handlers"
)

func registerAuthRoutes(mux *http.ServeMux, h *handlers.AuthHandler, authMW func(http.Handler) http.Handler) {
	mux.HandleFunc("POST /api/auth/login", h.Login)
	mux.Handle("GET /api/auth/me", authMW(http.HandlerFunc(h.Me)))
}
