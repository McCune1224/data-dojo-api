# Use the official Golang image as the base image
FROM golang:1.19-alpine

# Set your working directory
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY cmd/api/ api/ ./

# Change the working directory to the subdirectory containing main.go
WORKDIR /app/cmd/api

# Build the application
RUN go build -o main .

# Expose the port your application will run on
EXPOSE 8080

# Run the application
CMD ["/app/cmd/api/main"]
