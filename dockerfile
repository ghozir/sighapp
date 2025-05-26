# Stage 1: Builder
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go mod & download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary (tanpa debug info biar kecil)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o sighapp

# Stage 2: Runtime (super kecil)
FROM alpine:latest

WORKDIR /root/

# Tambahin cert untuk HTTPS request
RUN apk --no-cache add ca-certificates

# Copy hasil build
COPY --from=builder /app/sighapp .

# Expose port sesuai app
EXPOSE 3000

CMD ["./sighapp"]
