package app

import (
	"context"
	"net/http"

	v1 "github.com/desulaidovich/auth/auth/http/v1"
	"github.com/desulaidovich/auth/auth/middleware"
	"github.com/desulaidovich/auth/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type App struct {
	*config.Config
	*sqlx.DB
}

func NewApp(options ...func(*App)) *App {
	app := &App{}

	for _, call := range options {
		call(app)
	}
	return app
}

func WithConfig(cfg *config.Config) func(*App) {
	return func(a *App) {
		a.Config = cfg
	}
}

func WithDB(db *sqlx.DB) func(*App) {
	return func(a *App) {
		a.DB = db
	}
}

func (app *App) Run(ctx context.Context) error {
	mux := http.NewServeMux()

	s := &http.Server{
		Addr:    ":8080",
		Handler: middleware.Request(mux),
	}

	key := []byte(app.SecretKey)

	mux.HandleFunc("POST /api/v1/generate", v1.NewGenerateHandler(app.DB, key).GenerateToken)
	mux.HandleFunc("POST /api/v1/refresh", v1.NewRefreshHandler(app.DB, key).RefreshToken)

	return s.ListenAndServe()
}
