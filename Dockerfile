FROM golang:1.24.2 AS builder

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN curl -sSf https://atlasgo.sh | sh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

FROM golang:1.24.2 AS final

WORKDIR /app

COPY --from=builder /app /app
COPY --from=builder /go/bin/air /usr/local/bin/air
COPY --from=builder /usr/local/bin/atlas /usr/local/bin/atlas
COPY --from=builder /go/bin/swag /usr/local/bin/swag

CMD ["air"]
