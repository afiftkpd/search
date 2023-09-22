# Use the official Go 1.21 image as a parent image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the local source files to the container's working directory
COPY . .

RUN go clean

# Build the Go application
RUN GO_ENABLED=0 GOOS=linux go build -o main -installsuffix cgo -ldflags '-w'

# Expose port 8081
EXPOSE 8081

# Command to run the application
CMD ["./main"]