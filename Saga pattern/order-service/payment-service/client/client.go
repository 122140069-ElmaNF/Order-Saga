package client

import (
	"context"
	"log"
	"time"

	"saga-app/gen/paymentpb"

	"google.golang.org/grpc"
)

// PaymentService digunakan untuk komunikasi dengan PaymentService via gRPC
type PaymentService struct {
	api paymentpb.PaymentServiceClient
}

// Membuat instance baru dari PaymentService
func NewPaymentService(conn grpc.ClientConnInterface) *PaymentService {
	return &PaymentService{
		api: paymentpb.NewPaymentServiceClient(conn),
	}
}

// Kirim permintaan proses pembayaran berdasarkan ID order
func (p *PaymentService) Pay(orderID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := &paymentpb.ProcessPaymentRequest{
		OrderId: orderID,
	}

	response, err := p.api.ProcessPayment(ctx, request)
	if err != nil {
		return "", err
	}

	log.Printf("[PaymentService] Pembayaran diproses. Payment ID: %s, Status: %s", response.PaymentId, response.Status)
	return response.PaymentId, nil
}

// Kirim permintaan refund berdasarkan ID pembayaran
func (p *PaymentService) Refund(paymentID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := &paymentpb.RefundPaymentRequest{
		PaymentId: paymentID,
	}

	response, err := p.api.RefundPayment(ctx, request)
	if err != nil {
		return err
	}

	log.Printf("[PaymentService] Refund diproses. Payment ID: %s, Status: %s", paymentID, response.Status)
	return nil
}
