FROM golang:1.24.1 AS builder

RUN go install github.com/air-verse/air@latest
RUN curl -sSf https://atlasgo.sh | sh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

FROM golang:1.24.1 AS final

WORKDIR /app

COPY --from=builder /app /app
COPY --from=builder /go/bin/air /usr/local/bin/air

CMD ["air"]
