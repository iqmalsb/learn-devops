# --- STAGE 1: Build ---
FROM golang:1.25-alpine AS builder

# Install build tools for CGO (required for the SQLite driver)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy dependency files first (optimizes Docker layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o enterprise-app main.go

# --- STAGE 2: Run ---
# ... after the build stage ...
FROM alpine:latest

# 1. Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# 2. Create the data directory and give the new user ownership
RUN mkdir -p /app/database && chown -R appuser:appgroup /app

# 3. Copy the binary
COPY --from=builder /app/enterprise-app .
COPY ./seeder/payroll-dump.sql /app/database/seeder/payroll-dump.sql

# 4. Switch to the non-root user
USER appuser

EXPOSE 8080
CMD ["./enterprise-app"]