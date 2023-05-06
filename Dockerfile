# Use the offical Golang image to create a build artifact.
FROM golang:1.20-bullseye as builder

# Copy local code to the container image.
WORKDIR /go/app
COPY . .

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app main.go

# Use a Docker multi-stage build to create a lean production image.
FROM gcr.io/distroless/base-debian11
COPY --from=builder /go/app/ .

# Run the service binary.
CMD ["/app"]
