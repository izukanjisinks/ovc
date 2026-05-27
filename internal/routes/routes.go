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

	childRepo := repository.NewChildRepository(db)
	lookupRepo := repository.NewLookupRepository(db)
	reportRepo := repository.NewReportRepository(db)
	highlightRepo := repository.NewHighlightRepository(db)
	dashboardRepo := repository.NewDashboardRepository(db)

	childService := services.NewChildService(childRepo)
	lookupService := services.NewLookupService(lookupRepo)
	reportService := services.NewReportService(reportRepo)
	highlightService := services.NewHighlightService(highlightRepo)
	dashboardService := services.NewDashboardService(dashboardRepo)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(authService, userService)
	childHandler := handlers.NewChildHandler(childService)
	lookupHandler := handlers.NewLookupHandler(lookupService)
	reportHandler := handlers.NewReportHandler(reportService)
	highlightHandler := handlers.NewHighlightHandler(highlightService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	authMW := middleware.Auth(cfg, userRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Health)

	registerAuthRoutes(mux, authHandler, authMW)
	registerUserRoutes(mux, userHandler, authMW, permRepo)
	registerChildRoutes(mux, childHandler, authMW, permRepo)
	registerLookupRoutes(mux, lookupHandler, authMW, permRepo)
	registerReportRoutes(mux, reportHandler, authMW, permRepo)
	registerDashboardRoutes(mux, dashboardHandler, highlightHandler, authMW, permRepo)

	return middleware.CORS(mux)
}
