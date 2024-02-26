# Start from the official golang image
FROM golang:1.22-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
EXPOSE 8080

# Command to run the executable
RUN go build -o main ./cmd/htmx-go/main.go
CMD ["./main"]
