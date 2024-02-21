package main

import (
	"github.com/lallison21/to-do/internal/lib/logger/sl"
	"github.com/lallison21/to-do/internal/storage/postgres"
	"log/slog"
	"os"

	"github.com/lallison21/to-do/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting to-go", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := postgres.New(cfg.StorageConfig)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	id, err := storage.CreateRole("Администратор", 1)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	log.Info("created role", slog.Int64("id", id))

	id, err = storage.CreateRole("Администратор", 1)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	// TODO: init router: chi, "chi render"

	// TODO: run server:
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
