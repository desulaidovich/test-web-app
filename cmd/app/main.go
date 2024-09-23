package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/desulaidovich/auth/app"
	"github.com/desulaidovich/auth/config"
	"github.com/jmoiron/sqlx"
)

func main() {
	ctx := context.Background()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.ErrorContext(ctx, "Config", slog.String("Message", err.Error()))
	}

	conn, err := sqlx.ConnectContext(ctx, "pgx", cfg.DB)
	if err != nil {
		logger.ErrorContext(ctx, "Database", slog.String("Message", err.Error()))
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.ErrorContext(ctx, "Databse", slog.String("Message", err.Error()))
		}
	}()

	app := app.NewApp(
		app.WithConfig(cfg),
		app.WithDB(conn),
	)

	go func() {
		logger.InfoContext(ctx, "HTTP server listening...")

		if err := app.Run(ctx); err != nil {
			panic(err)
		}

		logger.InfoContext(ctx, "Stopped serving connections.")
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	logger.InfoContext(ctx, "Shutdown complete.")
}
