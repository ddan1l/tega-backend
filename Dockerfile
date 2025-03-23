FROM golang:1.24.1 AS builder

RUN go install github.com/air-verse/air@latest

WORKDIR /app/tega-backend

COPY air.toml ./
COPY go.mod ./

RUN go mod download 

RUN ls -la

COPY . .

CMD ["air"]