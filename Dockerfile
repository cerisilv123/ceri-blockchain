# syntax=docker/dockerfile:1

FROM golang:1.19

# Set working directory inside container
WORKDIR /app

# Copy go source code into container
COPY . .

# Build Go application and install dependencies (go get)
RUN go mod download

# Create a new image with a smaller base
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/ceri-blockchain /app/ceri-blockchain

# Expose the port your application listens on
EXPOSE 8080

# Command to run the application
CMD ["/app/my-golang-app"]