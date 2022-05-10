# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

ADD . /app/

RUN go mod download

RUN go build .

EXPOSE 4141

CMD ["/app/fhir"]