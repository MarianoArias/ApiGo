# Define base image
FROM golang:alpine as builder

# Update packages
RUN apk update 

# Install git (required for fetching the dependencies)
RUN apk add --no-cache git

# Create "app" directory
RUN mkdir /app

# Add source to "app" directory
ADD . /app/

# Set "app" directory as working directory
WORKDIR /app

# Build application
RUN go build -o customers cmd/customers/main.go

# Create user
RUN adduser -S -D -H -h /app appuser

# Set user
USER appuser

# Run application
CMD ["./customers"]