# go-agentic-template

A minimal, strict Go application starter for coding agents. This is the
Go-native counterpart of [rigerc/ts-agentic-template](https://github.com/rigerc/ts-agentic-template),
with the same goals: short agent instructions, testable application boundaries,
strict checks, reproducible CI, and no speculative infrastructure.

**This is a project template for AI-assisted Go development, not an LLM agent framework.**
The application uses only the Go standard library; no framework, ORM, or
runtime agent dependency is required.

## Requirements

- Go **1.26.10** (specified in `go.mod`; this line is supported by the pinned golangci-lint v2.10.1)
- GNU Make and a POSIX shell for the command shortcuts
- [`golangci-lint` builder](https://golangci-lint.run/docs/welcome/install/) to compile the pinned v2.10.1 custom binary with `anti-slop-go` v1.3.0
- Git and network access on the first `make lint` (the plugin and linter must be downloaded)
- Optional: [`govulncheck`](https://go.dev/doc/tutorial/govulncheck) for `make audit`

## Start

```sh
go mod download  # no third-party modules in the starter
make check
make run
```

The executable logs a test greeting to standard output using `log/slog`.
Development logs use key-value text; production and test logs use newline-delimited
JSON. Edit `cmd/app/main.go` and `internal/app/` to turn it into your application.

To set a different environment, export variables directly, or copy and load
`sample.env` explicitly:

```sh
cp sample.env .env
set -a; . ./.env; set +a
make run
```

The application does **not** load `.env` automatically. Keep `.env` untracked.

## Commands

| Command | Action |
| --- | --- |
| `make check` | Format check, vet, strict lint, tests with coverage gate, build |
| `make fix` / `make format` | Format Go source with `gofmt` |
| `make lint` | `go vet` + pinned custom golangci-lint with anti-slop-go |
| `make test` | All tests with race detector |
| `make test-coverage` | Race-enabled tests + **80%** statement coverage minimum |
| `make build` | Build `dist/app` |
| `make run` | Run the executable |
| `make audit` | Scan reachable vulnerabilities using `govulncheck` |
| `make hooks-install` | Enable optional local pre-commit format + lint checks |
| `make clean` | Delete generated binary and coverage report |

The coverage floor applies to `internal/...` (application logic), not to the
thin `cmd/app` composition root. Every package is built and tested separately
via `go test -race ./...`.

## Layout

```text
AGENTS.md               Single source of project instructions for coding agents
cmd/app/main.go         Composition root and process entry point
internal/app/           Example business logic with injected logging
internal/config/        Startup environment validation (no global env reads)
internal/logging/       slog logger factory (JSON production, text development)
.golangci.yml           Standard and anti-slop linting rules
.custom-gcl.yml         Reproducible custom linter build, with pinned plugin
.github/workflows/ci.yml  Checks, vulnerability review, and workflow security
.github/dependabot.yml  Weekly Go module and GitHub Actions updates
.githooks/pre-commit    Optional local Git hook
sample.env              Environment-variable contract
```

## Strictness and anti-slop

The anti-slop rules come from
[JacobJNilsson/anti-slop-go](https://github.com/JacobJNilsson/anti-slop-go),
not a homegrown port of the TypeScript Oxlint rules.
[`.custom-gcl.yml`](.custom-gcl.yml) pins **anti-slop-go v1.3.0** and
**golangci-lint v2.10.1**, the supported pairing documented by the plugin.
This uses Go 1.26 rather than Go 1.27 because golangci-lint v2.10.1 officially supports Go 1.26; its Go 1.27 support starts with v2.13.0.
The `plugin` import is required for registration. On the first `make lint`,
`golangci-lint custom` compiles `./custom-gcl`; subsequent checks reuse it
until the builder configuration changes. The generated binary is ignored by Git.

[`.golangci.yml`](.golangci.yml) runs all eight default anti-slop rules plus
`noanyparam`, `nointerfacereturn`, `justifypanic` and `fullstructcomp` (with a
three-field threshold). `errsemantics` is deliberately not enabled because
some configuration tests validate diagnostic message content. The same binary
also runs `errcheck`, `govet`, `ineffassign`, `staticcheck`, `unused`, `gocritic`,
`gosec`, `revive` and `nolintlint`.

To build and invoke the tool manually:

```sh
golangci-lint custom
./custom-gcl run ./...
```

A stock `golangci-lint run` **does not load the module plugin**; use `make lint`
or the generated binary. Prefer dependency injection, explicit errors,
concrete types, `context.Context` when cancellation is needed, and tests
that exercise real behavior rather than monkeypatching.

## CI and supply chain

GitHub Actions runs `make check` on pull requests and pushes to `main`.
Additional jobs cover new high-severity dependency advisories in PRs,
`govulncheck` on a weekly schedule, and workflow inspection via `actionlint`
and `zizmor`. A required aggregate `CI result` job rejects failing checks.
The Go module uses no external runtime dependencies; adding one should be
an explicit, reviewed decision. Dependabot updates Go modules and Actions
weekly with a 30-day version-update cooldown.

Tooling versions and GitHub Action commits are pinned in the workflow.
CI builds the same anti-slop-enabled custom golangci-lint binary as local checks.
The scheduled `govulncheck` report is non-blocking for existing advisories;
new high-severity dependency advisories are blocked on pull requests.

## Use as a new repository

1. Create a repository, preferably `rigerc/go-agentic-template`.
2. Copy this directory into it and initialize Git (`git init`).
3. If using a different repository path, update the module path in `go.mod`
   and all `github.com/rigerc/go-agentic-template` imports.
4. Optionally enable GitHub's **Template repository** setting.

`LICENSE` retains the BSD-2-Clause attribution from the original template.
