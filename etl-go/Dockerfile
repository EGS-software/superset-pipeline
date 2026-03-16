# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder
LABEL authors="jao"

WORKDIR /app

# Copy dependency files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o main ./main.go

# Stage 2: Final lightweight image
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# We don't copy the .env file because we will pass
# the variables via docker-compose for this test.
CMD ["./main"]