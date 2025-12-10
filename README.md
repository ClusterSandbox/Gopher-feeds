# Gopher-Feeds
Gopher-Feeds (a.k.a. WebScraper) is a small Go HTTP service that demonstrates a minimal JSON API and readiness/health endpoints.

The code uses the `net/http` standard library together with `chi` for routing, `cors` middleware for CORS handling and `godotenv` to load environment variables from a `.env` file.

## Features

- Minimal HTTP server with routing mounted under `/v1`
- Readiness/health endpoint (`GET /v1/healthz`)
- Example error endpoint (`GET /v1/err`) which returns a JSON error
- Small JSON helper utilities for consistent JSON responses

## Quick start

Prerequisites:
- Go 1.25+ installed
- Git (to clone the repo)

Clone and run:

```bash
    git clone https://github.com/ClusterSandbox/Gopher-feeds.git
    cd Gopher-feeds
    # create a .env file with a PORT value, for example:
    echo "PORT=8080" > .env
    go run .
```

Or build a binary and run:

```bash
    go build -o gopher-feeds .
    ./gopher-feeds
```

The server expects a `PORT` environment variable (loaded from `.env` if present). If `PORT` is missing the server will exit with a fatal error.

## Endpoints

- `GET /v1/healthz` — readiness/health check. Returns a 200 JSON response (currently an empty object `{}`).
- `GET /v1/err` — returns a structured error response (HTTP 400) via the helper `respondWithError`.

Examples with curl:

```bash
# Health check
curl -i http://localhost:8080/v1/healthz

# Error endpoint
curl -i http://localhost:8080/v1/err
```

Expected response shape for errors (JSON):

```json
{
	"error": "Something went wrong"
}
```

## Important files

- `main.go` — application entrypoint, router setup, CORS middleware and server start
- `handler_readiness.go` — readiness handler
- `handler.go` — example error handler
- `json.go` — helpers for JSON responses (`respondWithJSON`, `respondWithError`)

## Dependencies

Managed via Go modules. Key dependencies in `go.mod`:

- github.com/go-chi/chi — lightweight router
- github.com/go-chi/cors — CORS middleware
- github.com/joho/godotenv — load `.env` files into environment variables

Install dependencies (if needed):

```bash
go mod download
```

## Contributing

Feel free to open issues or PRs. Keep changes small and focused. If you add new endpoints, include simple curl examples and tests.


