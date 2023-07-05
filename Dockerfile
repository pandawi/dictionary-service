# Use the official Golang base image with the specified version
FROM golang:1.20.4 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o service

# Use a minimal base image to create the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/service .
COPY --from=builder /app/words_clean.txt .

RUN ls -la

# Expose the port that the application listens on
EXPOSE 8000

# Define the command to run your application
CMD ["./service"]
