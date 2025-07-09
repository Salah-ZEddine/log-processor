# --- Stage 1: Build binary ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build optimized static binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -a -installsuffix cgo \
    -ldflags='-w -s -extldflags "-static"' \
    -o log-processor ./cmd/main.go

RUN chmod +x log-processor

# --- Stage 2: Minimal runtime image ---
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --chmod=755 --from=builder /app/log-processor /usr/local/bin/log-processor

# Run it
CMD ["/usr/local/bin/log-processor"]
