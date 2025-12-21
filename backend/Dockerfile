# Use official Golang image as builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum (if exists)
COPY go.mod ./
# COPY go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main ./cmd/api/main.go

# Use a small alpine image for production
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./main"]
