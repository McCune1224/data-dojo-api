# Use an official Golang runtime as a parent image
FROM golang:1.19-alpine

# Set the working directory to /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the relevant directories to the container
COPY cmd/ /app/cmd/
COPY api/ /app/api/

# Build the application
RUN go build -o ./bin/main /app/cmd/api/main.go

# Expose the port your application will run on
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./bin/main"]
