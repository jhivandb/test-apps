# Status 500 Logger

A simple Go HTTP server that always returns HTTP status 500 (Internal Server Error) and logs every request it receives.

## What It Does

This application creates an HTTP server that:
- Listens on port 8080 (configurable via `PORT` environment variable)
- Responds to all requests with HTTP status 500
- Logs each request with timestamp, method, path, and remote address
- Returns "Internal Server Error" as the response body

## Use Cases

- Testing error handling in applications
- Simulating service failures
- Monitoring and alerting system testing
- Load testing error scenarios

## Running Locally

### With Go installed:

```bash
go run main.go
```

### With Docker:

Build the image:
```bash
docker build -t status-500-app .
```

Run the container:
```bash
docker run -p 8080:8080 status-500-app
```

## Configuration

Set the `PORT` environment variable to change the listening port:

```bash
PORT=3000 go run main.go
```

Or with Docker:
```bash
docker run -p 3000:3000 -e PORT=3000 status-500-app
```

## Testing

Send a request to the server:

```bash
curl http://localhost:8080
```

Expected output:
```
Internal Server Error
```

The server logs will show:
```
[2025-11-20 10:30:45.123] GET / 127.0.0.1:54321 - Returning 500
```

## Docker Multi-Stage Build

The Dockerfile uses a multi-stage build process:
1. **Build stage**: Uses `golang:1.21-alpine` to compile the Go application
2. **Final stage**: Uses `alpine:latest` with only the compiled binary, resulting in a minimal image size
