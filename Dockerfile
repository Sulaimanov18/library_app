FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o library_app ./cmd/main.go

FROM debian:bullseye-slim

WORKDIR /app

# Copy only the built binary from builder stage
COPY --from=builder /app/library_app .

# Set environment variable (optional)
ENV PORT=8080

# Start the app
CMD ["./library_app"]


