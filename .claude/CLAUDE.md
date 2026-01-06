# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`explain` is a CLI tool in the senzing-tools suite that explains Senzing SDK messages by parsing message IDs (e.g., `SZSDK60010032`) and opening the corresponding documentation page in a browser.

## Common Commands

```bash
# Build
make clean build          # Build binaries to target/<os>-<arch>/explain

# Test
make test                 # Run all tests with gotestfmt output
go test -v ./...          # Run all tests directly
go test -v ./explainer    # Run tests for specific package

# Lint
make lint                 # Run golangci-lint, govulncheck, and cspell

# Fix lint issues
make fix                  # Run all available fixers
make fix-gofumpt          # Format code with gofumpt

# Update dependencies
make dependencies         # Update and tidy go modules

# Development setup (one-time)
make dependencies-for-development
```

## Architecture

### Command Pattern with Explainer Interface

The codebase uses the [Command Pattern](https://en.wikipedia.org/wiki/Command_pattern) via the `Explainer` interface:

```go
type Explainer interface {
    Explain(ctx context.Context) error
}
```

Implementations:

- **MessageExplainer** - Parses `SZSDK` message IDs, maps component IDs to documentation pages, and opens browser
- **NullExplainer** - Null object pattern; returns nil when no message ID is provided

### Package Structure

- `cmd/` - Cobra command setup; `root.go` selects which Explainer to run based on CLI flags
- `explainer/` - Core logic: Explainer interface and implementations
- `makefiles/` - Platform-specific make targets (linux.mk, darwin.mk, windows.mk)

### CLI Framework

Uses `spf13/cobra` and `spf13/viper` with Senzing's `go-cmdhelping` library for standardized option handling. Options are defined as `option.ContextVariable` in `cmd/root.go`.

## Testing Patterns

- Tests use `testify/require` for assertions
- All tests call `test.Parallel()` at the start
- Use `test.Context()` to get context
- Test files use `_test` package suffix (e.g., `package explainer_test`)

## Linting

Strict golangci-lint configuration at `.github/linters/.golangci.yaml` with 100+ linters enabled. Code must pass before commits.

## Environment Variables

- `LD_LIBRARY_PATH=/opt/senzing/er/lib` - Required when using Senzing libraries
- `SENZING_TOOLS_MESSAGE_ID` - Alternative to `--message-id` flag
