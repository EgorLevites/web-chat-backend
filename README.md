
# Web Chat Backend

This is the backend service for the Web Chat application. It is implemented in Go and handles WebSocket connections, API endpoints, and broadcasting chat messages.

## Features

- Manages WebSocket connections.
- Provides REST API endpoints for:
  - Returning the number of connected clients (`/api/clients`).
  - Sending a welcome message (`/api/welcome`).
- Broadcasts chat messages to all connected clients.

## Requirements

- Go 1.20+
- Gorilla WebSocket library (`github.com/gorilla/websocket`)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/your-repo/web-chat-backend.git
   cd web-chat-backend
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Run the server:
   ```
   go run main.go
   ```

4. The backend server will start on `http://localhost:8080`.

## API Endpoints

- `GET /api/clients`: Returns the current number of connected clients.
- `GET /api/welcome`: Returns a welcome message.

## WebSocket

- WebSocket connection endpoint: `/ws`
- Messages must be JSON formatted:
  ```json
  {
    "username": "YourUsername",
    "content": "YourMessage"
  }
  ```

## Dockerfile Explanation

The backend application is containerized using a multi-stage Dockerfile:

### Dockerfile

```dockerfile
# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder

# Install necessary packages
RUN apk update && apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o web-chat-backend main.go

# Stage 2: Create a minimal image to run the application
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/web-chat-backend .

# Expose the backend port
EXPOSE 8080

# Run the backend application
CMD ["./web-chat-backend"]
```

### Steps to Build and Run Docker Container

1. Build the Docker image:
   ```bash
   docker build -t web-chat-backend .
   ```

2. Run the container:
   ```bash
   docker run -d -p 8080:8080 --name web-chat-backend-container web-chat-backend
   ```

## Docker Compose

The backend is configured to work with Docker Compose along with the frontend:

### docker-compose.yml

```yaml
version: '3.8'

services:
  frontend:
    build:
      context: ./web-chat-frontend
    ports:
      - "80:80"
    depends_on:
      - backend

  backend:
    build:
      context: ./web-chat-backend
    ports:
      - "8080:8080"
```

To start both the backend and frontend together:
```bash
docker-compose up --build
```

Access the frontend at [http://localhost](http://localhost), and the backend API at [http://localhost:8080](http://localhost:8080).

## CORS

- All origins are allowed for development purposes.
- Modify the `enableCORS` function in `main.go` to restrict origins for production.

## License

MIT License
