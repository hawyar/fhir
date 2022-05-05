# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY /server/*.go /app/server

COPY go.mod /app
COPY go.sum /app

COPY /*.go /app

COPY Caddyfile /app/Caddyfile

COPY *.go /app

RUN go mod download

RUN go build -o server

EXPOSE 4141

CMD ["/server"]