package config

import (
    "ecommerce-app/pkg/logger"
    "github.com/joho/godotenv"
    "os"
)

// Config holds application configuration
type Config struct {
    Port        string
    DatabaseURL string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
    log := logger.New()

    // Load .env file if present (optional)
    err := godotenv.Load()
    if err != nil {
        log.Info("No .env file found, using system environment variables")
    }

    cfg := &Config{
        Port:        getEnv("PORT", "8080"),
        DatabaseURL: getEnv("DATABASE_URL", "host=localhost user=postgres password=root dbname=ecommerce_db port=5432 sslmode=disable"),
    }

    log.Info("Configuration loaded successfully")
    return cfg, nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}