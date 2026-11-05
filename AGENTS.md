# Go template

## Setup

Use Go 1.26.10 (as specified in `go.mod`) and the custom `golangci-lint` v2.10.1 with `anti-slop-go` v1.3.0.
Application packages use only the Go standard library. Do not add dependencies
without a specific reason.

## Verification

Before opening a pull request, run:

```sh
make check
```

`make check` enforces `gofmt`, `go vet`, anti-slop Go analysis and standard lint, race-enabled tests,
an 80% statement coverage floor on business packages, and a compiled binary.
Use `make fix` for formatting. Do not use plain `golangci-lint run`: it lacks
the anti-slop module plugin. `make lint` builds and runs `./custom-gcl` from
`.custom-gcl.yml` with the pinned plugin. Run `make audit` with `govulncheck` before
dependency upgrades or releases.

## Structure

- `cmd/app/` is the executable and composition root.
- `internal/app/` holds business logic, with colocated `_test.go` files.
- `internal/config/` validates environment settings once at startup.
- `internal/logging/` constructs injected `slog` loggers.
- `sample.env` documents the keys validated by `internal/config/`.
- `.custom-gcl.yml` pins the custom linter and anti-slop module.
- `.golangci.yml` configures Go linters and the enabled anti-slop rules.

## Code style

Return errors rather than panicking. Keep I/O at boundaries and inject loggers,
writers, and environment lookups; avoid global mutable state. Use `slog` instead
of ad hoc print logging. Add tests for invalid input and failure paths. Update
`sample.env` and configuration validation together.

## Boundaries

Keep changes scoped to the task. Ask first before adding deployment,
publication, or infrastructure automation. Never commit `.env`, credentials,
coverage output, or binaries. Do not bypass checks to make CI green.
