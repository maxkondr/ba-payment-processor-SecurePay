#!/bin/sh
set -e

go get -u github.com/googleapis/googleapis || true
go get -u -d github.com/maxkondr/ba-payment-processor-dispatcher || true

# build payment-processor-dispatcher Golang stubs
protoc --proto_path=$GOPATH/src/github.com/googleapis/googleapis \
    --proto_path=$GOPATH/src/github.com/maxkondr/ba-payment-processor-dispatcher/proto \
    --go_out=plugins=grpc:$GOPATH/src/github.com/maxkondr/ba-payment-processor-dispatcher/proto \
    $GOPATH/src/github.com/maxkondr/ba-payment-processor-dispatcher/proto/processor-dispatcher.proto

protoc --proto_path=$GOPATH/src/github.com/googleapis/googleapis \
    --proto_path=$GOPATH/src/github.com/maxkondr/ba-payment-processor-secure-pay/proto \
    --proto_path=$GOPATH/src/ \
    --go_out=plugins=grpc:$GOPATH/src/github.com/maxkondr/ba-payment-processor-secure-pay/proto \
    $GOPATH/src/github.com/maxkondr/ba-payment-processor-secure-pay/proto/SecurePay.proto

CGO_ENABLED=0 GOOS=linux go build -a -o ba-payment-processor-secure-pay .
# docker build -t maxkondr/ba-payment-processor-secure-pay .
# docker push maxkondr/ba-payment-processor-secure-pay
