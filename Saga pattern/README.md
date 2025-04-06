# Tugas 2 Pemrograman Web Lanjut  
## Nama: Elma Nurul Fatika  | 122140069

---

## Deskripsi

Proyek ini mengimplementasikan **Saga Pattern** dalam arsitektur **microservices** untuk menangani transaksi terdistribusi menggunakan:

- **Bahasa Pemrograman**: Golang (Go)
- **Komunikasi Antar Service**: gRPC
---

## Layanan Microservice

### 1. Order Service
- Endpoint: `CreateOrder`, `CancelOrder`
- Status: `PENDING`, `COMPLETED`, `CANCELLED`

### 2. Payment Service
- Endpoint: `ProcessPayment`, `RefundPayment`
- Status: `SUCCESS`, `FAILED`, `REFUNDED`

### 3. Shipping Service
- Endpoint: `StartShipping`, `CancelShipping`
- Status: `PENDING`, `SHIPPED`, `CANCELLED`

---

## Alur Saga Orchestration

1. Orchestrator memanggil `CreateOrder`
2. Jika sukses, dilanjutkan ke `ProcessPayment`
3. Jika sukses, dilanjutkan ke `StartShipping`
4. Jika gagal di tahap shipping:
   - Jalankan kompensasi: `CancelShipping` → `RefundPayment` → `CancelOrder`
5. Jika gagal di tahap pembayaran:
   - Jalankan kompensasi: `CancelOrder`


## Cara Menjalankan

### 1. Jalankan ketiga service:

```bash
go run order-service/server/main.go
go run payment-service/server/main.go
go run shipping-service/server/main.go
