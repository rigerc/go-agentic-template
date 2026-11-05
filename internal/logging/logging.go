// Package logging constructs the application's structured logger.
package logging

import (
	"io"
	"log/slog"

	"github.com/rigerc/go-agentic-template/internal/config"
)

// New returns a logger that writes human-readable text locally and JSON in
// non-development environments. The writer is injected to make logging testable.
func New(out io.Writer, cfg config.Config) *slog.Logger {
	options := &slog.HandlerOptions{Level: cfg.LogLevel}
	if cfg.Environment == config.Development {
		return slog.New(slog.NewTextHandler(out, options))
	}

	return slog.New(slog.NewJSONHandler(out, options))
}
