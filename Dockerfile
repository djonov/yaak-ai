# Start from the official Golang image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy whole local directory to the container
COPY . .

# Download all dependencies
RUN go mod download

# Build the Go app
RUN go build -o server ./cmd/server

# Expose port 9595 to the outside world
EXPOSE 9595

# Command to run the executable
CMD ["./server"]