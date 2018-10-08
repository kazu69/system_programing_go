FROM golang:alpine

RUN apk add --update \
    openssl-dev \
    pcre-dev \
    git \
    gcc \
    musl-dev \
    linux-headers \
    sudo

WORKDIR /go/src/github.com/kazu69/system_programing_go/chapter17/
ADD container.go ./container.go
ADD rootfs ./rootfs

RUN go get github.com/opencontainers/runc/libcontainer
RUN go get golang.org/x/sys/unix
RUN go build -o container container.go
