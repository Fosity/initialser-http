FROM golang:1.5.1

MAINTAINER Leonlau "Leonlau@aliyun.com"

ADD . $GOPATH/src/initialser-http
ADD image $GOPATH/src/golang.org/x/image
RUN go install golang.org/x/image
RUN go get github.com/leonlau/initialser-http
RUN go install -a initialser-http

EXPOSE 80
CMD ["initialser-http"]