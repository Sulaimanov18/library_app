FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o library_app ./cmd/main.go

FROM debian:bullseye-slim

WORKDIR /app

# Install PostgreSQL client tools and goose
RUN apt-get update && \
    apt-get install -y postgresql-client curl && \
    curl -L https://github.com/pressly/goose/releases/download/v3.18.0/goose_linux_x86_64 -o /usr/local/bin/goose && \
    chmod +x /usr/local/bin/goose && \
    rm -rf /var/lib/apt/lists/*

# Copy only the built binary from builder stage
COPY --from=builder /app/library_app .

# Set environment variable (optional)
ENV PORT=8080

# Start the app
CMD ["./library_app"]


