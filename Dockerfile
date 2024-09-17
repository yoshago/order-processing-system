# Use the official Golang image as the base image
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY cmd cmd
COPY internal internal

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy && go get ./...

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/server

# Use a minimal image as the base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder container
COPY --from=builder /app/main .

# Copy .env file
COPY .env .

# Install PostgreSQL client for database interaction
RUN apk add --no-cache postgresql-client

# Command to run the executable
CMD ["/root/main"]
