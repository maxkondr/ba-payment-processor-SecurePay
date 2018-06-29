package main

import (
	"log"
	"net"

	"github.com/maxkondr/ba-payment-processor-A/proto"
	"github.com/maxkondr/ba-payment-processor-A/server"
	"google.golang.org/grpc"
)

// main start a gRPC server and waits for connection
func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	paymentProcessorA.RegisterPaymentProcessorAServer(grpcServer, &paymentProcessorAImpl.Server{})
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
