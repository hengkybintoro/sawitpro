# Dockerfile definition for Backend application service.

# From which image we want to build. This is basically our environment.
FROM golang:1.21-alpine as Build

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary Go mod and sum files
# COPY go.mod .
# COPY go.sum .

# Download dependencies using go modules
# RUN go mod download

# This will copy all the files in our repo to the inside the container at root location.
COPY . .

# Build our binary at root location.
RUN GOPATH= go build -o /main cmd/main.go

####################################################################
# This is the actual image that we will be using in production.
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# We need to copy the binary from the build image to the production image.
COPY --from=Build /main .

# This is the port that our application will be listening on.
EXPOSE 1323

# This is the command that will be executed when the container is started.
ENTRYPOINT ["./main"]