package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/malinatrash/url-shortener/internal/config"
	"github.com/malinatrash/url-shortener/internal/http-server/middleware/logger"
	"github.com/malinatrash/url-shortener/internal/lib/logger/handlers/slogpretty"
	"github.com/malinatrash/url-shortener/internal/lib/logger/sl"
	"github.com/malinatrash/url-shortener/internal/storage/sqlite"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	router := chi.NewRouter()
	log.Info("starting server")
	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// TODO: run server

}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "production"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
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

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
