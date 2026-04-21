package routes

import (
	"database/sql"
	"net/http"

	"github.com/izukanji/ovc/internal/config"
	"github.com/izukanji/ovc/internal/handlers"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/repository"
	"github.com/izukanji/ovc/internal/services"
)

func Setup(db *sql.DB, cfg *config.Config) http.Handler {
	userRepo := repository.NewUserRepository(db)
	permRepo := repository.NewPermissionRepository(db)

	authService := services.NewAuthService(userRepo, permRepo, cfg)
	userService := services.NewUserService(userRepo, permRepo)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(authService, userService)

	authMW := middleware.Auth(cfg, userRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Health)

	// Auth
	mux.HandleFunc("POST /api/auth/login", authHandler.Login)
	mux.Handle("GET /api/auth/me", authMW(http.HandlerFunc(authHandler.Me)))

	// Users — auth + permission per route
	mux.Handle("POST /api/users", authMW(
		middleware.RequirePermission(permRepo, "user-management", "create")(
			http.HandlerFunc(userHandler.Create),
		),
	))
	mux.Handle("GET /api/users", authMW(
		middleware.RequirePermission(permRepo, "user-management", "read")(
			http.HandlerFunc(userHandler.List),
		),
	))
	mux.Handle("PUT /api/users/{id}", authMW(
		middleware.RequirePermission(permRepo, "user-management", "update")(
			http.HandlerFunc(userHandler.Update),
		),
	))
	mux.Handle("DELETE /api/users/{id}", authMW(
		middleware.RequirePermission(permRepo, "user-management", "delete")(
			http.HandlerFunc(userHandler.Delete),
		),
	))

	return middleware.CORS(mux)
}
