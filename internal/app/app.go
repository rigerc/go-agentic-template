// Package app contains business logic that is independent of process state.
package app

import (
	"log/slog"

	"github.com/rigerc/go-agentic-template/internal/config"
)

// Greet returns a greeting for the supplied name.
func Greet(name string) string {
	return "Hello, " + name + "!"
}

// Run uses an injected logger and validated configuration, making the
// application behavior testable without global state or module-level mocks.
func Run(logger *slog.Logger, cfg config.Config, name string) string {
	greeting := Greet(name)
	logger.Info("greeting generated", "app_env", cfg.Environment, "greeting", greeting)

	return greeting
}
