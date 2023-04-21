# Use the official Golang image as the base image
FROM golang:1.19-alpine

# Set your working directory
WORKDIR /app

# Get the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project to the working directory
COPY cmd/ api/ ./
# Build the application
RUN go build -o ./bin/main ./cmd/api/main.go

# Expose the port your application will run on
EXPOSE 8080

# Run the application
CMD ["./bin/main"]
