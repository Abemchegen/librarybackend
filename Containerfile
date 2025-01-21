# Use the official Golang image as the base
FROM golang:1.23.4

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files first (to install dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the Go application
RUN go build -o main .

COPY .env .env
COPY config ./config

# Expose the port your Go app runs on
EXPOSE 5000

# Run the app
CMD ["./main"]
