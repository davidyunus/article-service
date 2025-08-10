FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install dependencies first for layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN go build -o article-service ./cmd/server

# Runtime image
FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/article-service .
EXPOSE 8080

CMD ["./article-service"]
