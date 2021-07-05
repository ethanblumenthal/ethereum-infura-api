# Start from golang base image
FROM golang:alpine as builder

# Enable go modules
ENV GO111MODULE=on

# Install git (alpine image does not have git in it)
RUN apk update && apk add --no-cache git

# Set current working directory
WORKDIR /app

# Copy go mod and sum files to cache dependencies
COPY go.mod ./
COPY go.sum ./

# Download all dependencies if not in cache
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# CGO_ENABLED is disabled for cross system compilation
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

# Finally our multi-stage to build a small image
# Start a new stage from scratch
FROM scratch

# Copy the Pre-built binary file
COPY --from=builder /app/bin/main .

# Run executable
CMD ["./main"]