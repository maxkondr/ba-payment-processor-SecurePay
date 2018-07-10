package paymentProcessorSecurePayImpl

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/maxkondr/ba-payment-processor-dispatcher/proto"

	context "golang.org/x/net/context"
)

var (
	paymentProcessorInfo = &PaymentProcessorDispatcher.PaymentProcessorInfo{
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

// Server implementation for SecurePay
type Server struct{}

// GetInfo is interface func ba-payment-processor-A.GetInfo
func (s *Server) GetInfo(ctx context.Context, in *empty.Empty) (*PaymentProcessorDispatcher.PaymentProcessorInfo, error) {
	return paymentProcessorInfo, nil
}

// Pay is interface func ba-payment-processor-A.Pay
func (s *Server) Pay(context context.Context, req *PaymentProcessorDispatcher.MakePaymentRequest) (*PaymentProcessorDispatcher.MakePaymentResponse, error) {
	return &PaymentProcessorDispatcher.MakePaymentResponse{
		Uuid:         req.Uuid,
		Success:      true,
		Errstring:    "",
		Md5:          "md5_hash",
		AvsCode:      "AvsCode",
		CavvResponse: "CavvResponse",
		Cvv2Response: "Cvv2Response",
	}, nil
}
