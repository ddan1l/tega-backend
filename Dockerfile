FROM golang:1.24.1 AS builder

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

CMD ["air"]