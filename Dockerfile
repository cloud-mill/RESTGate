FROM golang:1.19 as builder

WORKDIR /go/src/catache.com/RESTGate/
COPY . .
RUN go get
RUN go build -v -ldflags "-w -s -X main.Version=1.0.0 -X main.Build=`date +%FT%T%z`" -o bin/RESTGate-linux-amd64

FROM debian:buster-slim

MAINTAINER catache.com

ARG RESTATE_CONFIG_PATH
ENV RESTATE_CONFIG_PATH ${RESTATE_CONFIG_PATH}

RUN apt-get update \
    && apt-get install -y ca-certificates tzdata \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /application

COPY --from=builder /go/src/catache.com/RESTGate/bin/RESTGate-linux-amd64 .

ENTRYPOINT ./RESTGate-linux-amd64
