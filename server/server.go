package paymentProcessorAImpl

import (
	"fmt"

	"github.com/maxkondr/ba-payment-processor-A/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// Server implementation for ba-payment-processor-A
type Server struct{}

// Pay is interface func ba-payment-processor-A.Pay
func (s *Server) Pay(context context.Context, req *paymentProcessorA.MakePaymentRequest) (*paymentProcessorA.MakePaymentResponse, error) {
	headers, _ := metadata.FromIncomingContext(context)
	fmt.Println("YAKUT incoming MD: ", headers)

	return &paymentProcessorA.MakePaymentResponse{Uuid: req.GetPayment().Uuid,
		Success: true, Message: ""}, nil
}
