#!/bin/sh
set -e

protoc --proto_path=$GOPATH/src/github.com/maxkondr/ba-payment-processor-A/proto \
    --go_out=plugins=grpc:$GOPATH/src/github.com/maxkondr/ba-payment-processor-A/proto \
    $GOPATH/src/github.com/maxkondr/ba-payment-processor-A/proto/processor-A.proto

go build .