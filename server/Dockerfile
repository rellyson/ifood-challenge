## production base image (targeted as base)
FROM golang:1.18-alpine as base

WORKDIR /opt/ifood/server

RUN apk add --no-cache git build-base

COPY . .

## download deps and build golang bin
RUN go mod download
RUN go build -o bin/server main.go

## run binary
CMD ["./bin/server"]

#############################################

## development image (targeted as develop)
FROM golang:1.18-alpine as develop

WORKDIR /opt/ifood/server

RUN apk add --no-cache autoconf git build-base \
    automake libtool gettext gettext-dev make g++ texinfo curl

# Setup fswatch
RUN cd /root && wget https://github.com/emcrisostomo/fswatch/releases/download/1.16.0/fswatch-1.16.0.tar.gz && \
    tar -xf fswatch-1.16.0.tar.gz && rm fswatch-1.16.0.tar.gz
RUN cd /root/fswatch-1.16.0 && ./configure && make && make install