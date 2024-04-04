FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .


RUN go build -o /app/bin/go-docker /app/cmd/go-docker/main.go

CMD ["/app/bin/go-docker"]