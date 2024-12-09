# Stage 1: Build the Go application
FROM golang:1.22.1 AS go-builder

# Set the working directory for Go
WORKDIR /app

COPY server/ ./

# Download dependencies
# RUN go mod download # No dependencies needed for this app

# Build the Go application
RUN go build -o main ./main.go

# Stage 2: Build the Node.js application
FROM node:18 AS node-builder

# Set the working directory for Node.js
WORKDIR /test-engine

# Copy Node.js package files and install dependencies
COPY test-engine/package.json ./
RUN npm install

# Install Jasmine globally
RUN npm install -g jasmine

# Copy the rest of the test-engine directory
COPY test-engine/ ./

# Stage 3: Create the final runtime image
FROM golang:1.22.1

# Set the working directory for the final image
WORKDIR /app

# Copy the built Go application from the Go builder stage
COPY --from=go-builder /app/main ./main

# Copy static files and templates for the Go app
COPY server/static/ ./static/

# Copy Node.js test engine from the Node builder stage
WORKDIR /test-engine
COPY --from=node-builder /test-engine/ ./

# Expose the server port
EXPOSE 5130

WORKDIR /app

# Command to start the server
CMD ["./main"]
