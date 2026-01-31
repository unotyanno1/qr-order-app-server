package qrcode

import (
	"net/http"

	"github.com/labstack/echo/v4"

	qrcodedomain "github.com/unotyanno1/qr-order-app-server/domain/qrcode"
	qrcodeusecase "github.com/unotyanno1/qr-order-app-server/usecase/qrcode"

)

// Handler handles HTTP requests for QR code operations
type Handler struct {
	useCase *qrcodeusecase.UseCase
}

// NewHandler creates a new QR code handler
func NewHandler(useCase *qrcodeusecase.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// GetQRCode handles GET /qr_code/:seat_number
func (h *Handler) GetQRCode(c echo.Context) error {
	seatNumber := c.Param("seat_number")

	req := &qrcodedomain.QRCodeRequest{
		SeatNumber: seatNumber,
	}

	result, err := h.useCase.GetQRCode(req);
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"result": result,
	})
}
