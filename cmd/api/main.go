package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/izukanji/ovc/internal/config"
	"github.com/izukanji/ovc/internal/database"
	"github.com/izukanji/ovc/internal/routes"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if err := godotenv.Load(); err != nil {
		logger.Info("no .env file found, using environment variables")
	}

	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("database connection failed", zap.Error(err))
	}
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		logger.Fatal("migrations failed", zap.Error(err))
	}

	handler := routes.Setup(db, cfg)

	logger.Info("server starting", zap.String("port", cfg.Port))
	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		logger.Fatal("server failed", zap.Error(err))
	}
}
