package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/maxkondr/ba-payment-processor-secure-pay/server"
	"github.com/maxkondr/ba-proto/paymentProcessor"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go-opentracing"
	zipkinot "github.com/openzipkin/zipkin-go-opentracing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	zipkinURL       = flag.String("zipkinUrl", "http://zipkin.tracing:9411/api/v1/spans", "Zipkin server URL")
	zipkinTracer    opentracing.Tracer
	zipkinCollector zipkintracer.Collector
	myPort          = 7777
)

func initTracerZipkin() {
	var err error
	zipkinCollector, err = zipkinot.NewHTTPCollector(*zipkinURL)
	if err != nil {
		grpclog.Error("err", err)
		os.Exit(1)
	}

	myName := os.Getenv("MY_POD_NAME")
	myIP := os.Getenv("MY_POD_IP")
	myNS := os.Getenv("MY_POD_NAMESPACE")

	var (
		debug       = false
		hostPort    = fmt.Sprintf("%s:%d", myIP, myPort)
		serviceName = fmt.Sprintf("%s.%s(%s)", myName, myNS, myIP)
	)

	recorder := zipkinot.NewRecorder(zipkinCollector, debug, hostPort, serviceName)
	zipkinTracer, err = zipkinot.NewTracer(recorder)
	if err != nil {
		grpclog.Error("err", err)
		os.Exit(1)
	}
	opentracing.SetGlobalTracer(zipkinTracer)
}

func init() {
	initTracerZipkin()
}

// main start a gRPC server and waits for connection
func main() {
	flag.Parse()

	grpcLog := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)

	defer zipkinCollector.Close()

	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", myPort))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer(grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor()),
	),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_opentracing.UnaryServerInterceptor()),
		),
	)

	paymentProcessor.RegisterPaymentProcessorServer(grpcServer, &paymentProcessorSecurePayImpl.Server{})
	// start the server
	grpclog.Infof("Start listening on port %d", myPort)
	if err := grpcServer.Serve(lis); err != nil {
		grpclog.Fatalf("failed to serve: %s", err)
	}
}
