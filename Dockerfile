FROM golang:1.22.1

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY ./api/go.mod ./api/go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY ./api .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /ciri2-pc-microservice

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 6000

# Run
CMD ["/ciri2-pc-microservice"]