# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY config.json ./

RUN go mod download

COPY *.go ./

RUN go build -o /fhir

EXPOSE 4141

CMD ["/fhir"]