package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Config holds application configuration loaded from environment variables.
// No credentials are hardcoded; every value comes from the environment.
type Config struct {
	Port string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

// Load reads configuration from environment variables, applying defaults
// where a variable is not set.
func Load() Config {
	cfg := Config{
		Port: getEnv("PORT", "8080"),

		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "warehouse"),
		DBPassword: getEnv("DB_PASSWORD", "warehouse"),
		DBName:    getEnv("DB_NAME", "warehouse"),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
	}

	// Validate the port is numeric so the server fails fast with a clear
	// message instead of a cryptic listen error later.
	if _, err := strconv.Atoi(cfg.Port); err != nil {
		log.Fatalf("invalid PORT %q: must be a number", cfg.Port)
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

// DSN builds the PostgreSQL connection string.
func (c Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}
