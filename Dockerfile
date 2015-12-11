FROM golang:alpine

# Installing the postgres client on the transporter container
RUN apk update
RUN apk add postgresql-client

RUN apk update
RUN apk add git

# Installing files

RUN mkdir -p /go/src/github.com/wawandco/transporter

RUN go get github.com/lib/pq
RUN go get github.com/go-sql-driver/mysql

WORKDIR /go/src/github.com/wawandco/transporter
