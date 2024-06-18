# Start from a base golang image
FROM golang:1.21

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire contents of the current directory to the working directory
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main.go

# Command to run the executable
CMD ["./main"]
