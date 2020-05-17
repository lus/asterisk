FROM golang:1.14-alpine

# Define the directory we should work in
WORKDIR /app

# Download the neccessary go modules
COPY go.mod go.sum
RUN go mod download

# Build the application and define the starting command
COPY . .
RUN go build -o asterisk .
CMD ["./asterisk"]