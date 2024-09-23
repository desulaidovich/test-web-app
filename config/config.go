package config

import (
	"github.com/desulaidovich/auth/internal/env"
)

type Config struct {
	DB        string `env:"DATABASE_URL"`
	SecretKey string `env:"SECRET_KEY"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Read(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
