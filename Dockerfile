# syntax=docker/dockerfile:1
FROM ubuntu:latest
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git

FROM golang:1.19-alpine

WORKDIR /app

#Copy and run files for dependencies
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./



FROM golang:1.19-alpine

RUN go get github.com/qronica/core/migrations
#build go
RUN go build -o qronica .

EXPOSE 8080

RUN ./qronica serve
#CMD [ "/qronica" ]