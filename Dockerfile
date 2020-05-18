FROM golang:1.14-alpine

# Define the directory we should work in
WORKDIR /app

# Download the neccessary go modules
COPY go.mod go.sum
RUN go mod download

# Build the application and define the starting command
COPY . .
RUN go build \
        -o asterisk \
        -ldflags "\
            -X github.com/Lukaesebrot/asterisk/static.Mode=prod \
            -X github.com/Lukaesebrot/asterisk/static.Version=$(git describe --tags --abbrev=0)+$(git describe --tags | sed -n 's/^[0-9]\+\.[0-9]\+\.[0-9]\+-\([0-9]\+\)-.*$/\1/p')" \
        .
CMD ["./asterisk"]