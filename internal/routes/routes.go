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

	// Repositories
	userRepo := repository.NewUserRepository(db)

	// Services
	authService := services.NewAuthService(userRepo, cfg)
	userService := services.NewUserService(userRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(authService, userService)

	r.GET("/health", handlers.Health)

	api := r.Group("/api")

	registerAuthRoutes(api, authHandler, cfg)
	registerUserRoutes(api, userHandler, cfg)
}
