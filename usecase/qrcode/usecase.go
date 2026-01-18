package qrcode

import (
	"log"

	"github.com/unotyanno1/qr-order-app-server/domain/qrcode"

)

// UseCase handles QR code business logic
type UseCase struct{}

// NewUseCase creates a new QR code use case
func NewUseCase() *UseCase {
	return &UseCase{}
}

// GetQRCode processes the QR code request and logs the seat number
func (uc *UseCase) GetQRCode(req *qrcode.QRCodeRequest) error {
	log.Printf("Seat number: %s", req.SeatNumber)
	return nil
}
