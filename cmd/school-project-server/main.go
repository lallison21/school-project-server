package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lallison21/school-project-server/internal/http-server/handlers/url/role"
	mwLogger "github.com/lallison21/school-project-server/internal/http-server/middleware"
	"github.com/lallison21/school-project-server/internal/lib/logger/sl"
	"github.com/lallison21/school-project-server/internal/storage/postgres"
	pretty_slog "github.com/lallison21/school-project-server/pkg/prettySlog"
	"log/slog"
	"net/http"
	"os"

	"github.com/lallison21/school-project-server/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//application.New().Run()

	cfg := config.MustLoad()
	config.New()

	log := setupLogger(cfg.Env)

	log.Info("starting school-project-server", slog.String("env", cfg.Env))

	storage, err := postgres.New(cfg.StorageConfig)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/get_roles", role.GetRoles(log, storage))
	router.Get("/get_role/{id}", role.GetRoleById(log, storage))
	router.Post("/create_role", role.CreateRole(log, storage))

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			pretty_slog.NewHandler(&slog.HandlerOptions{Level: slog.LevelDebug}),
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
