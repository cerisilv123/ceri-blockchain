# syntax=docker/dockerfile:1

FROM golang:1.19

# Set the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the project
RUN CGO_ENABLED=0 GOOS=linux go build -o /ceri-blockchain cmd/main.go

# Expose the port
EXPOSE 8080

# Run the application
CMD ["/ceri-blockchain"]




