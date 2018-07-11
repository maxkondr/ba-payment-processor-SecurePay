#!/bin/sh
set -e

go get -u github.com/googleapis/googleapis || true
go get -u -d github.com/maxkondr/ba-proto || true

# build payment-processor Golang stubs
protoc --proto_path=$GOPATH/src/github.com/googleapis/googleapis \
    --proto_path=$GOPATH/src/github.com/maxkondr/ba-proto \
    --go_out=plugins=grpc:$GOPATH/src/ \
    $GOPATH/src/github.com/maxkondr/ba-proto/paymentProcessor/payment-processor.proto

CGO_ENABLED=0 GOOS=linux go build -a -o ba-pp-SecurePay .
# docker build -t maxkondr/ba-pp-secure-pay .
# docker push maxkondr/ba-pp-secure-pay
