FROM golang:1.5.1

MAINTAINER Leonlau "Leonlau@aliyun.com"

ADD . $GOPATH/src/initialser-http
ADD . $GOPATH/src/golang.org/x/image
RUN cd $GOPATH/src/golang.org/x
RUN git clone git@github.com:golang/image.git
RUN go install image
RUN cd -
RUN go get github.com/leonlau/initialser-http
RUN go install -a initialser-http

EXPOSE 80
CMD ["initialser-http"]