#!/bin/sh
set -e

# go get -u github.com/maxkondr/ba-payment-processor-dispatcher

protoc --proto_path=/Porta/repos/yakut_github/porta-microservices/googleapi/googleapis \
    --proto_path=$GOPATH/src/github.com/maxkondr/ba-payment-processor-A/proto \
    --proto_path=$GOPATH/src/ \
    --go_out=plugins=grpc:$GOPATH/src/github.com/maxkondr/ba-payment-processor-A/proto \
    $GOPATH/src/github.com/maxkondr/ba-payment-processor-A/proto/SecurePay.proto

CGO_ENABLED=0 GOOS=linux go build -a -o ba-payment-processor-secure-pay .
# docker build -t maxkondr/ba-payment-processor-secure-pay .
# docker push maxkondr/ba-payment-processor-secure-pay
