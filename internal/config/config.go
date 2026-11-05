// Package config validates application configuration at process startup.
package config

import (
	"errors"
	"fmt"
	"log/slog"
)

// Environment is a validated application execution environment.
type Environment string

const (
	// Development is the default local environment.
	Development Environment = "development"
	// Test is the environment for automated tests.
	Test Environment = "test"
	// Production is the production environment.
	Production Environment = "production"
)

// Config contains validated application settings.
type Config struct {
	Environment Environment
	LogLevel    slog.Level
}

// LookupFunc provides environment values without coupling validation to os.Environ.
type LookupFunc func(string) (string, bool)

// Load validates environment values. Defaults apply only when a key is absent.
// Invalid settings return errors naming the offending variable.
func Load(lookup LookupFunc) (Config, error) {
	if lookup == nil {
		return Config{}, errors.New("environment lookup must not be nil")
	}

	cfg := Config{
		Environment: Development,
		LogLevel:    slog.LevelInfo,
	}

	if raw, ok := lookup("APP_ENV"); ok {
		switch Environment(raw) {
		case Development, Test, Production:
			cfg.Environment = Environment(raw)
		default:
			return Config{}, fmt.Errorf("APP_ENV: must be development, test, or production; got %q", raw)
		}
	}

	if raw, ok := lookup("LOG_LEVEL"); ok {
		switch raw {
		case "debug":
			cfg.LogLevel = slog.LevelDebug
		case "info":
			cfg.LogLevel = slog.LevelInfo
		case "warn":
			cfg.LogLevel = slog.LevelWarn
		case "error":
			cfg.LogLevel = slog.LevelError
		default:
			return Config{}, fmt.Errorf("LOG_LEVEL: must be debug, info, warn, or error; got %q", raw)
		}
	}

	return cfg, nil
}
