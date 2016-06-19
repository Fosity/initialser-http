FROM golang:1.5.1

MAINTAINER Leonlau "Leonlau@aliyun.com"

ADD . $GOPATH/src/app
RUN go get app
RUN CGO_ENABLED=0 go install -a app

EXPOSE 80
CMD app