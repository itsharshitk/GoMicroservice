# Use the official Golang image as the base image
FROM golang:1.21
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./
# Install dependencies
RUN go mod download
# Copy the source code
COPY . .
# Build the Go app
RUN go build -o main .
# Expose port 8081 to the outside world
EXPOSE 8000
# Command to run the executable
CMD ["./main"]