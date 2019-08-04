package paymentProcessorSecurePayImpl

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/maxkondr/ba-proto/paymentProcessor"
	"github.com/sirupsen/logrus"

	context "golang.org/x/net/context"
)

var (
	paymentProcessorInfo = &paymentProcessor.PaymentProcessorInfo{
		IOnlinePaymentProcessor: 1,
		Processor:               "SecurePay",
		WebLink:                 "",
		Handler:                 "",
		Callback:                "",
		ExtAuth:                 false,
		Obsolete:                false,
		Remittance:              false,
		PostProcessing:          false,
		EmailAuth:               false,
		RemoteCcStorage:         false,
	}
)

func getLogger(context context.Context) *logrus.Entry {
	return ctxlogrus.Extract(context)
}

// Server implementation for SecurePay
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

// GetInfo is interface func ba-payment-processor-A.GetInfo
func (s *Server) GetInfo(ctx context.Context, in *empty.Empty) (*paymentProcessor.PaymentProcessorInfo, error) {
	getLogger(ctx).Info("Received request")
	return paymentProcessorInfo, nil
}

// Pay is interface func ba-payment-processor-A.Pay
func (s *Server) Pay(context context.Context, req *paymentProcessor.MakePaymentRequest) (*paymentProcessor.MakePaymentResponse, error) {
	getLogger(context).Info("Received request with uuid:", req.GetUuid())
	return &paymentProcessor.MakePaymentResponse{
		Uuid:         req.GetUuid(),
		Success:      true,
		Errstring:    "",
		Md5:          "md5_hash",
		AvsCode:      "AvsCode",
		CavvResponse: "CavvResponse",
		Cvv2Response: "Cvv2Response",
	}, nil
}
