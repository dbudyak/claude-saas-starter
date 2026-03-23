// Package config loads application configuration from environment variables.
package config

import (
	"log/slog"
	"os"
)

// Config holds all application configuration.
type Config struct {
	DatabaseURL string
	JWTSecret   string
	Port        string
	AppURL      string
}

// Load reads configuration from environment variables.
// Required variables cause a fatal log if missing.
func Load() *Config {
	return &Config{
		DatabaseURL: required("DATABASE_URL"),
		JWTSecret:   required("JWT_SECRET"),
		Port:        optional("PORT", "8080"),
		AppURL:      optional("APP_URL", "http://localhost"),
	}
}

func required(key string) string {
	v := os.Getenv(key)
	if v == "" {
		slog.Error("required environment variable not set", "key", key)
		os.Exit(1)
	}
	return v
}

func optional(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
