# syntax=docker/dockerfile:1
FROM golang:1.24.3-alpine

# Install git (needed for downloading Go modules)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files first (for better caching)
COPY go.mod go.sum ./

# Copy the entire project
COPY . .

# Build the application
RUN go build -o main server.go

# Expose the port your app runs on
EXPOSE 3000

# Start the app
CMD ["./main"]
