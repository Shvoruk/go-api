FROM golang:1.24.1-alpine AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . .

# Build
RUN go build -o go-api ./cmd/main.go

FROM alpine:latest

# App directory
WORKDIR /app

# Copy built binary
COPY --from=builder /app/go-api .

# Expose port
EXPOSE 8080

# Start
ENTRYPOINT ["./go-api"]