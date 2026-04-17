package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izukanji/ovc/internal/config"
	"github.com/izukanji/ovc/internal/handlers"
	"github.com/izukanji/ovc/internal/middleware"
	"github.com/izukanji/ovc/internal/repository"
	"github.com/izukanji/ovc/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(r *gin.Engine, db *pgxpool.Pool, cfg *config.Config) {
	r.Use(middleware.CORS())

	userRepo := repository.NewUserRepository(db)
	permRepo := repository.NewPermissionRepository(db)

	authService := services.NewAuthService(userRepo, permRepo, cfg)
	userService := services.NewUserService(userRepo, permRepo)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(authService, userService)

	r.GET("/health", handlers.Health)

	api := r.Group("/api")
	authMW := middleware.Auth(cfg, userRepo)

	registerAuthRoutes(api, authHandler, authMW)
	registerUserRoutes(api, userHandler, authMW, permRepo)
}
