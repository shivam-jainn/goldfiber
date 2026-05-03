# GoldFiber

A starter template for a Go web service using Fiber, PostgreSQL, and Docker Compose.

## Overview

This repository demonstrates a clean layout for a Go + Fiber application with:

- Fiber web server in `cmd/server/main.go`
- environment-based configuration in `internal/config`
- PostgreSQL connection via `pgxpool` in `internal/database`
- request logging middleware in `internal/logger`
- route registration in `internal/handler`
- Docker Compose support for local development

## Repo Structure

- `cmd/server/main.go` — application bootstrap and graceful shutdown
- `internal/app/app.go` — Fiber app setup and middleware registration
- `internal/config` — env loading, defaults, and required field validation
- `internal/database` — database pool creation and health check
- `internal/handler` — HTTP route handlers
- `internal/logger` — logging interface and Fiber logging middleware
- `infra/` — Compose files for development and production
- `db/` — SQL schema and queries
- `Makefile` — helper commands for development tasks

## Getting Started

### Prerequisites

- Go 1.25+
- PostgreSQL for local development
- Docker & Docker Compose for infra-based workflows
- `air` for hot reloading (optional, used by `make dev`)

### Run locally

1. Copy or create `.env` with required variables:

```env
PORT=8080
DEBUG=false
ENV=dev
LOG_LEVEL=info
DB_URL=postgres://user:password@localhost:5432/goldfiber?sslmode=disable
```

2. Start the app:

```bash
go run ./cmd/server
```

3. Alternatively, use the Makefile dev command:

```bash
make dev
```

### Docker Compose

Start development infrastructure:

```bash
make dev-infra
```

Stop it:

```bash
make dev-infra-down
```

Reset infra:

```bash
make dev-infra-reset
```

## Environment Variables

The application reads configuration from environment variables via `internal/config`.

- `PORT` — HTTP server port (default `8080`)
- `DEBUG` — debug mode flag (default `false`)
- `ENV` — environment name (default `dev`)
- `LOG_LEVEL` — log verbosity (default `info`)
- `DB_URL` — PostgreSQL connection URL (required)

## Health Checks

- `GET /health` — simple health check
- `GET /health/details` — JSON health response

## Development Notes

- Use `internal/app` to register routes and middleware.
- Keep business logic in `internal/service` or `internal/repository` as the app grows.
- The current logger abstraction supports swapping logger implementations.

## Extend the Template

Recommended next improvements:

- add API versioning (`/api/v1`)
- add auth routes and middleware
- add tests for config loading and route handlers
- add CI lint/test pipeline
- add structured error handling and validation

## License

MIT
