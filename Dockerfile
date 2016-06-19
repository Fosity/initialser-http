FROM golang:1.5.1

MAINTAINER Leonlau "leonlau@aliyun.com"

ADD . $GOPATH/src/initialser-http

RUN go get github.com/leonlau/initialser-http
RUN go install -a initialser-http

EXPOSE 80
CMD initialser-http http -d $GOPATH/src/initialser-http/resource -p 80