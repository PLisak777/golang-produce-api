# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT=8000

RUN go build 

CMD ["./golang-produce-api"]
