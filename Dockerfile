FROM golang:1.17-alpine as builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o bot

FROM alpine:3.15

RUN apk add --update bash openssl

WORKDIR /app

COPY --from=builder /app/bot .
COPY /public /app

CMD ["./bot --help"]