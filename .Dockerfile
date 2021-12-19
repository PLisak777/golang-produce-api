# syntax=docker/dockerfile:1

FROM golang:1.17.5-alpine
WORKDIR /go/golang-produce-api/
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go get -d -v github.com/gorilla/mux
EXPOSE 4000
CMD [ "/golang-produce-api" ]
