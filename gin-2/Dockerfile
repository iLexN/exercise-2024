# Use the official Go image as the base image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Set necessary environment variables and build your project.
ENV CGO_ENABLED=0 GIN_MODE=release

# Build the Go application
RUN go build -o api-app ./cmd/api


# FROM scratch
FROM golang:1.22-alpine

WORKDIR /app

COPY --from=builder /app/api-app .

# Set entry point.
ENTRYPOINT ["app/api-app"]

# CMD ["./api-app"]
