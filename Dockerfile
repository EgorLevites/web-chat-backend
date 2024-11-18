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