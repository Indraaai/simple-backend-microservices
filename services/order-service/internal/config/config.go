package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OrderDatabaseUrl string
}

func Load() (*Config, error) {
	// Load .env if it exists.
	// In production, environment variables can be provided
	// directly by Docker/Kubernetes/etc.
	_ = godotenv.Load()

	databaseURL := os.Getenv("ORDER_DATABASE_URL")

	if databaseURL == "" {
		return nil, fmt.Errorf("ORDER_DATABASE_URL is required")
	}

	return &Config{
		OrderDatabaseUrl: databaseURL,
	}, nil
}
