FROM golang:alpine
WORKDIR /go/src/github.com/shebesabrina/qs_go
ADD . /go/src/github.com/shebesabrina/qs_go
RUN go get
RUN go build
