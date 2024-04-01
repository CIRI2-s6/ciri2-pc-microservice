# Use a smaller base image for the builder stage
FROM golang:1.22.1 AS builder

# Set destination for COPY
WORKDIR /app


# Download Go modules
COPY ./api/go.mod ./api/go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY ./api .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ciri2-pc-microservice ./cmd

# Use a minimal base image for the final image
FROM scratch
# Copy the built binary from the builder stage
COPY --from=builder /app/ciri2-pc-microservice /ciri2-pc-microservice
COPY ./api/.env .

# Expose the port
EXPOSE 6000

# Run the binary
CMD ["/ciri2-pc-microservice"]
