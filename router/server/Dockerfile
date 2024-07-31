# Use the official Golang image with the specified version as the base image
FROM golang:1.22.5 

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container, excluding kube directory and liblab.config.json
COPY . .

# Remove the kube directory and liblab.config.json file from the copied source
RUN rm -rf kube liblab.config.json

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /main


# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/main"]


