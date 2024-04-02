FROM golang:latest  AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

ENV PORT=":1337"

EXPOSE 1337

RUN go build -o /app/bin/go-docker /app/main.go

CMD ["/app/bin/go-docker"]