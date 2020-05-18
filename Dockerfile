FROM golang:1.14-alpine

# Install git for the version string
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Define the directory we should work in
WORKDIR /app

# Download the neccessary go modules
COPY go.mod go.sum ./
RUN go mod download

# Build the application and define the starting command
COPY . .
RUN go build \
        -o asterisk \
        -ldflags "\
            -X github.com/Lukaesebrot/asterisk/static.Mode=prod \
            -X github.com/Lukaesebrot/asterisk/static.Version=$(git rev-parse --abbrev-ref HEAD)-$(git describe --tags --abbrev=0)-$(git log --pretty=format:'%h' -n 1)" \
        .
CMD ["./asterisk"]