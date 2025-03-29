FROM golang:1.24-alpine

# project dir
WORKDIR /app

# copy go.mod and go.sum to app dir
COPY go.mod go.sum ./
RUN go mod tidy

# copy all other folders and files
COPY . .

CMD ["go", "run", "cmd/api-server/main.go"]

# Exposing port
EXPOSE 8080