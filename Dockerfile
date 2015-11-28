FROM golang:alpine

# Installing the postgres client on the transporter container
RUN apk update
RUN apk add postgresql-client

# Installing files
RUN mkdir -p /go/src/github.com/wawandco/transporter
ADD . /go/src/github.com/wawandco/transporter
WORKDIR /go/src/github.com/wawandco/transporter
