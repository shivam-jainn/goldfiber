# GoldFiber

A starter template for a Go web service using Fiber, PostgreSQL, and Docker Compose.

## Overview

GoldFiber demonstrates a clean layout for a Go + Fiber application with:

- Fiber web server in `cmd/server/main.go`
- environment-based configuration in `internal/config`
- PostgreSQL connection management via `pgxpool` in `internal/database`
- request logging middleware in `internal/logger`
- route and handler registration in `internal/app` and `internal/handler`
- Docker Compose support for local development and observability
- SQL code generation and migration tooling in `db/` and `Makefile`

## Repo Structure

- `cmd/server/main.go` — application bootstrap and graceful shutdown
- `internal/app/app.go` — Fiber app setup, middleware, and route registration
- `internal/config` — environment loading, defaults, and required validation
- `internal/database` — database pool creation, ping health checks, and helpers
- `internal/handler` — HTTP route handlers
- `internal/logger` — logger abstraction, context helpers, and middleware
- `internal/repository` / `internal/service` — separation points for business logic
- `infra/` — Docker Compose manifest and observability configs
- `db/` — SQL schema, migrations, and generated query code
- `Makefile` — common development and infrastructure commands

## Getting Started

### Prerequisites

- Go 1.25+
- PostgreSQL for local development
- Docker & Docker Compose for infra-based workflows
- `air` for hot reloading (optional when using `make dev`)

### Environment

Create a `.env` file with the required values:

```env
PORT=8080
DEBUG=false
ENV=dev
LOG_LEVEL=info
DB_URL=postgres://user:password@localhost:5432/goldfiber?sslmode=disable
```

### Run Locally

Start the application directly:

```bash
go run ./cmd/server
```

Or run in development mode with hot reload:

```bash
make dev
```

### Docker Compose

Start development infrastructure:

```bash
make dev-infra
```

Stop development infrastructure:

```bash
make dev-infra-down
```

Reset development infrastructure and recreate containers:

```bash
make dev-infra-reset
```

Start development infrastructure with observability tooling:

```bash
make dev-infra-observability
```

## Useful Make Targets

- `make dev` — start the app with `air`
- `make lint` — run Go linters via `golangci-lint`
- `make test` — run Go unit tests
- `make dev-infra` — start local Docker Compose environment
- `make dev-infra-down` — stop local Docker Compose environment
- `make dev-infra-reset` — stop and recreate local Docker Compose environment
- `make dev-infra-observability` — start infra with observability services
- `make sqlc` — generate Go code from SQL queries with `sqlc`
- `make migrate` — run database migrations via `goose`

## Environment Variables

The application reads configuration from environment variables in `internal/config`.

- `PORT` — HTTP server port (default: `8080`)
- `DEBUG` — enable debug mode (default: `false`)
- `ENV` — runtime environment name (default: `dev`)
- `LOG_LEVEL` — log verbosity (default: `info`)
- `DB_URL` — PostgreSQL connection URL (required)

## API / Health Checks

- `GET /health` — basic health check
- `GET /health/details` — detailed JSON health response

## Development Notes

- Keep HTTP routing and middleware registration inside `internal/app`
- Place business logic in `internal/service` and persistence concerns in `internal/repository`
- Use `internal/logger` for structured logging and request context propagation

## License

MIT
