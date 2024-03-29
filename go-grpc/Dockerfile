# Use the official Go image as the base image
FROM golang:latest AS builder

# add linter
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# add protobuf related
RUN apt-get update \
    && apt-get install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN go build -o myapp

# Use a minimal Alpine-based image as the base image for the final container
#FROM alpine:latest
#
# Set the working directory inside the container
#WORKDIR /app

# Copy the built executable from the builder stage
#COPY --from=builder /app/myapp .

# Expose the port that the application listens on (if applicable)
EXPOSE 50051

# Set the command to run the executable
# CMD ["./myapp"]
CMD tail -f /dev/null