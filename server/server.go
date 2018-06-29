package paymentProcessorAImpl

import (
	"fmt"

	"github.com/maxkondr/ba-payment-processor-A/proto"
	context "golang.org/x/net/context"
)

// Server implementation for ba-payment-processor-A
type Server struct{}

// Pay is interface func ba-payment-processor-A.Pay
func (s *Server) Pay(context context.Context, req *paymentProcessorA.MakePaymentRequest) (*paymentProcessorA.MakePaymentResponse, error) {
	fmt.Println("Got payment request: ", req)

	// time.Sleep(3 * time.Second)

	return &paymentProcessorA.MakePaymentResponse{Uuid: req.GetPayment().Uuid,
		Success: true, Message: ""}, nil

}
