# Stage 1: Build Stage
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Modules files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files haven't changed.
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./Delivery/main.go

# Stage 2: Run Stage
FROM alpine:3.18

# Set working directory inside the container
WORKDIR /app

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Copy .env and other necessary files such as config


# Expose the port your app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]