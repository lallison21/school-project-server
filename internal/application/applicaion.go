package application

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/lallison21/school-project-server/internal/config"
	mwLogger "github.com/lallison21/school-project-server/internal/http-server/middleware"
	pretty_slog "github.com/lallison21/school-project-server/pkg/prettySlog"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Application struct {
	router *chi.Mux
	server *http.Server
	log    *slog.Logger
	db     *sqlx.DB
}

func New() *Application {
	cfg := config.New()

	log := setupLogger(cfg.Env)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	server := &http.Server{
		Addr:         cfg.Addr,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &Application{
		router: router,
		server: server,
		log:    log,
	}
}

func (app *Application) Run() {
	ctx := context.Background()

	_ = ctx

	if err := app.server.ListenAndServe(); err != nil {
		app.log.Error("failed to start server")
	}

	app.log.Error("server stopped")
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
