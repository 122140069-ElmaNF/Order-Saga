package main

import (
	"log"

	orderclient "saga-app/order-service/client"
	paymentclient "saga-app/payment-service/client"
	shippingclient "saga-app/shipping-service/client"

	"google.golang.org/grpc"
)

func main() {
	// Koneksi ke semua service
	orderConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer orderConn.Close()
	order := orderclient.NewOrderService(orderConn)

	paymentConn, _ := grpc.Dial("localhost:50052", grpc.WithInsecure())
	defer paymentConn.Close()
	payment := paymentclient.NewPaymentService(paymentConn)

	shippingConn, _ := grpc.Dial("localhost:50053", grpc.WithInsecure())
	defer shippingConn.Close()
	shipping := shippingclient.NewShippingClient(shippingConn)

	//test: Shipping Gagal → Semua Dikompensasi
	log.Println("Shipping gagal")
	testShippingFail(order, payment, shipping)

	// test: Payment Gagal → Batalkan Order
	log.Println("\n Payment gagal")
	testPaymentFail(order, payment)

	//test: Semua Berhasil
	log.Println("\n Berhasil!")
	testSuccess(order, payment, shipping)
}

//testcase

func testShippingFail(order *orderclient.OrderService, payment *paymentclient.PaymentService, shipping *shippingclient.ShippingClient) {
	orderID, err := order.PlaceOrder("Elmow", "fail-shipping")
	if err != nil {
		log.Println("Gagal membuat order:", err)
		return
	}
	paymentID, err := payment.Pay(orderID)
	if err != nil {
		log.Println("Gagal memproses pembayaran:", err)
		order.RevokeOrder(orderID)
		return
	}
	shippingID, err := shipping.StartShipping(orderID)
	if err != nil {
		log.Println("GAGAL mengirim:", err)
		log.Println("Kompensasi: Cancel Shipping → Refund → Cancel Order")
		shipping.CancelShipping(shippingID)
		payment.Refund(paymentID)
		order.RevokeOrder(orderID)
		return
	}
	log.Println("SUKSES shipping:", shippingID)
}

func testPaymentFail(order *orderclient.OrderService, payment *paymentclient.PaymentService) {
	orderID, err := order.PlaceOrder("Elmow", "fail-payment")
	if err != nil {
		log.Println("Gagal membuat order:", err)
		return
	}
	_, err = payment.Pay(orderID)
	if err != nil {
		log.Println("Gagal membayar:", err)
		log.Println("Kompensasi: Batalkan Order")
		order.RevokeOrder(orderID)
		return
	}
	log.Println("Pembayaran berhasil")
}

func testSuccess(order *orderclient.OrderService, payment *paymentclient.PaymentService, shipping *shippingclient.ShippingClient) {
	orderID, err := order.PlaceOrder("Elmow", "laptop")
	if err != nil {
		log.Println("Gagal membuat order:", err)
		return
	}
	paymentID, err := payment.Pay(orderID)
	if err != nil {
		log.Println("Gagal melakukan pembayaran:", err)
		order.RevokeOrder(orderID)
		return
	}
	shippingID, err := shipping.StartShipping(orderID)
	if err != nil {
		log.Println("Gagal mengirim barang:", err)
		payment.Refund(paymentID)
		order.RevokeOrder(orderID)
		return
	}
	log.Println("Order selesai dan berhasil dikirim!:", shippingID)
}
