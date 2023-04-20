# Stage 1: Build the Go binary
FROM golang:1.16 AS builder

# Set the working directory
WORKDIR /src

# Copy the Go source files into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o chatter-cli

# Stage 2: Create a minimal Docker image
FROM scratch

# Copy the chatter-cli binary from the builder stage to the minimal image
COPY --from=builder /src/chatter-cli /chatter-cli

# Set the entrypoint to run the chatter-cli binary
ENTRYPOINT ["/chatter-cli"]
