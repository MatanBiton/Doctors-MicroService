# syntax=docker/dockerfile:1

# Base image with Go and Alpine
FROM golang:1.23-alpine AS base

# Install dependencies only when needed
FROM base AS deps

# Install git for Go modules and version control
RUN apk add --no-cache git

# Set the working directory for the application
WORKDIR /app

# Copy protobuf files
COPY ./doctors_protobuf ./doctors_protobuf

# Copy go.mod and go.sum for dependency management
COPY ./server/go.mod ./server/go.mod
COPY ./server/go.sum ./server/go.sum

WORKDIR /app/server

# Build arguments for GitHub authentication (private repos)
ARG GITHUB_ACTOR
ARG GITHUB_TOKEN

# Configure Git to use token for accessing private repos
RUN git config --global url."https://${GITHUB_ACTOR}:${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

# Download Go modules
RUN go mod download

# Rebuild the source code only when needed
FROM base AS builder

# Set working directory and copy Go workspace
WORKDIR /app
COPY --from=deps $GOPATH $GOPATH

# Copy all source files
COPY . .

WORKDIR /app/server

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /doctors-ms

# Final production image
FROM golang:1.22-alpine

# Copy the compiled binary to the production image
COPY --from=builder /doctors-ms /doctors-ms

# Optionally copy the entire app directory
COPY --from=builder /app /app

# Expose the application port
EXPOSE 9090

# Set the command to run the binary
CMD ["/doctors-ms"]