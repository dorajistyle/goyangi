# Use the offical Golang image to create a build artifact.
FROM golang:1.20-alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

# RUN apt-get update && apt-get upgrade
# RUN apt-get -y install git make bash

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Set PKG_CONFIG_PATH for libvips
# ENV PKG_CONFIG_PATH=/usr/lib/x86_64-linux-gnu/pkgconfig
# RUN pkg-config --libs vips

COPY . .

# Build the command inside the container.
# RUN CGO_ENABLED=0 GOOS=linux go build -v -o app main.go # with this option, I got 'executor failed' error
RUN go build -v -o goyangi-backend main.go

# Distribution
FROM golang:1.20-alpine as prod

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

# EXPOSE 9090

COPY --from=builder /app /app

ENTRYPOINT ["/app/goyangi-backend"]

# Use a Docker multi-stage build to create a lean production image.
# FROM gcr.io/distroless/base-debian11

# RUN apt-get update && apt-get upgrade && \
#     apt-get install -y libvips-dev --no-install-recommends \ 
#     mkdir /app 

# COPY --from=builder /app/ .

# Run the service binary.
# CMD ["/app"]
