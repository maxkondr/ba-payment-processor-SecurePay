#!/bin/sh
set -e

protoc --proto_path=$GOPATH/src/github.com/maxkondr/ba-payment-processor-A/proto \
    --go_out=plugins=grpc:$GOPATH/src/github.com/maxkondr/ba-payment-processor-A/proto \
    $GOPATH/src/github.com/maxkondr/ba-payment-processor-A/proto/processor-A.proto

CGO_ENABLED=0 GOOS=linux go build -a .
# docker build -t maxkondr/ba-payment-processor-a .
# docker push maxkondr/ba-payment-processor-a