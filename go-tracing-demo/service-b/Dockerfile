# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the dependencies and tidy the go.mod file
RUN go mod tidy && go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o service-b .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/service-b .

# Expose the port that the application will listen on
EXPOSE 8081

# Set the environment variables for DataDog
ENV DD_SERVICE=service-b
ENV DD_ENV=local

# Start the application
CMD ["./service-b"]