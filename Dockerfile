# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application for the CLI
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goai ./cmd/cli

# Stage 2: Create the final, lightweight image
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/goai .

# Copy the .env file (or template)
# In a real scenario, you might manage secrets differently


# Set the entrypoint
ENTRYPOINT ["./goai"]

# The default command can be overridden
CMD ["--help"]
