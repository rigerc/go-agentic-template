package main

import (
	"fmt"
	"io"
	"os"

	"github.com/rigerc/go-agentic-template/internal/app"
	"github.com/rigerc/go-agentic-template/internal/config"
	"github.com/rigerc/go-agentic-template/internal/logging"
)

func main() {
	if err := run(os.Stdout, os.LookupEnv); err != nil {
		// The logger cannot be initialized before configuration is validated.
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(out io.Writer, lookup config.LookupFunc) error {
	cfg, err := config.Load(lookup)
	if err != nil {
		return err
	}

	logger := logging.New(out, cfg)
	app.Run(logger, cfg, "world")

	return nil
}
