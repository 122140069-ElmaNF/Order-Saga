package client

import (
	"context"
	"log"
	"time"

	"saga-app/gen/orderpb"

	"google.golang.org/grpc"
)

// Client untuk berinteraksi dengan OrderService gRPC
type OrderService struct {
	api orderpb.OrderServiceClient
}

// Inisialisasi client baru dengan koneksi gRPC
func NewOrderService(conn grpc.ClientConnInterface) *OrderService {
	return &OrderService{
		api: orderpb.NewOrderServiceClient(conn),
	}
}

// Buat order baru berdasarkan user ID dan item ID
func (o *OrderService) PlaceOrder(userID, itemID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := &orderpb.CreateOrderRequest{
		UserId: userID,
		ItemId: itemID,
	}

	response, err := o.api.CreateOrder(ctx, request)
	if err != nil {
		return "", err
	}

	log.Printf("[OrderService] Order berhasil dibuat. ID: %s, Status: %s", response.OrderId, response.Status)
	return response.OrderId, nil
}

// Batalkan order berdasarkan ID yang diberikan
func (o *OrderService) RevokeOrder(orderID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := &orderpb.CancelOrderRequest{
		OrderId: orderID,
	}

	response, err := o.api.CancelOrder(ctx, request)
	if err != nil {
		return err
	}

	log.Printf("[OrderService] Order dibatalkan. ID: %s, Status: %s", orderID, response.Status)
	return nil
}
