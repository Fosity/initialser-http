FROM golang:1.5.1

MAINTAINER Leonlau "Leonlau@aliyun.com"

ADD . $GOPATH/src/initialser-http

RUN go get github.com/leonlau/initialser-http
RUN go install -a initialser-http

EXPOSE 80
