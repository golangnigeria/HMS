# syntax=docker/dockerfile:1

# -----------------------------
# 1. Build Stage
# -----------------------------
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your project files
COPY . .

# Build the Go app from cmd/web
RUN go build -tags netgo -ldflags="-s -w" -o /app/app ./cmd/web

# -----------------------------
# 2. Runtime Stage
# -----------------------------
FROM alpine:latest

# Working directory in final image
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/app .

# Expose the app’s port
EXPOSE 8080

# Run the app
CMD ["./app"]
