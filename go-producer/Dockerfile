FROM golang:alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o bin/main ./cmd