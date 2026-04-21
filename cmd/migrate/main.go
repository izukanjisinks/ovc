package main

import (
	"flag"

	"github.com/joho/godotenv"
	"github.com/izukanji/ovc/internal/config"
	"github.com/izukanji/ovc/internal/database"
	"go.uber.org/zap"
)

func main() {
	command := flag.String("command", "up", "Migration command: up, down, reset, status")
	flag.Parse()

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

	switch *command {
	case "up":
		if err := database.RunMigrations(db); err != nil {
			logger.Fatal("migrate up failed", zap.Error(err))
		}
	case "down":
		if err := database.RollbackMigration(db); err != nil {
			logger.Fatal("migrate down failed", zap.Error(err))
		}
	case "reset":
		if err := database.ResetDatabase(db); err != nil {
			logger.Fatal("migrate reset failed", zap.Error(err))
		}
	case "status":
		if err := database.MigrationStatus(db); err != nil {
			logger.Fatal("migrate status failed", zap.Error(err))
		}
	default:
		logger.Fatal("unknown command", zap.String("command", *command))
	}
}
