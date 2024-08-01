# Use the official Golang image with the specified version as the base image
FROM golang:1.22.4-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container, excluding kube directory and liblab.config.json file from the copied source
COPY . .

# Remove the kube directory and liblab.config.json file
RUN rm -rf kube liblab.config.json

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Add a debugging step to check if the executable is present
RUN ls -la /app

# Command to run the executable
CMD ["./main"]

