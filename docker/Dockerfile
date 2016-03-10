FROM golang:1.5.1

RUN go get github.com/constabulary/gb/...

WORKDIR /go/src/app

ENV GOPATH=/go:/go/src/app:/go/src/app/vendor
