package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/maxkondr/ba-payment-processor-secure-pay/server"
	"github.com/maxkondr/ba-proto/paymentProcessor"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go-opentracing"
	zipkinot "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

var (
	zipkinURL       = flag.String("zipkinUrl", "http://zipkin.tracing:9411/api/v1/spans", "Zipkin server URL")
	zipkinTracer    opentracing.Tracer
	zipkinCollector zipkintracer.Collector
	myPort          = 7777
	logrusEntry     logrus.Entry
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

func initLogger() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.Formatter = &logrus.JSONFormatter{DisableTimestamp: true}
	log.SetOutput(l.Writer())

	logrusEntry = *logrus.NewEntry(l).WithFields(logrus.Fields{
		"process": os.Getenv("PORTA_APP_NAME"),
		"version": os.Getenv("PORTA_GIT_TAG") + ":" + os.Getenv("PORTA_GIT_COMMIT"),
	})
	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(&logrusEntry)
}

func grpcRequestIDPropagatorUnaryServerInterceptor(entry *logrus.Entry) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		headers, _ := metadata.FromIncomingContext(ctx)

		reqIDs, ok := headers["x-b3-traceid"]
		var reqID string
		if ok {
			reqID = reqIDs[0]
		}

		newCtx := ctxlogrus.ToContext(ctx, entry.WithFields(logrus.Fields{"request_id": reqID}))
		resp, err := handler(newCtx, req)
		return resp, err
	}
}

func init() {
	initTracerZipkin()
	initLogger()
}

// main start a gRPC server and waits for connection
func main() {
	flag.Parse()

	defer zipkinCollector.Close()

	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", myPort))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer(grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(&logrusEntry),
		),
	),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_opentracing.UnaryServerInterceptor(),
				grpcRequestIDPropagatorUnaryServerInterceptor(&logrusEntry),
				grpc_logrus.UnaryServerInterceptor(&logrusEntry)),
		),
	)

	paymentProcessor.RegisterPaymentProcessorServer(grpcServer, &paymentProcessorSecurePayImpl.Server{})
	// start the server
	grpclog.Infof("Start listening on port %d", myPort)
	if err := grpcServer.Serve(lis); err != nil {
		grpclog.Fatalf("failed to serve: %s", err)
	}
}
