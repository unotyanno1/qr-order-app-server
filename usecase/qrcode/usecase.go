package qrcode

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/unotyanno1/qr-order-app-server/domain/qrcode"
)

// ErrInvalidSeatNumber は座席番号が不正な形式の場合に返される。
var ErrInvalidSeatNumber = errors.New("invalid seat number")

// ErrSeatNotFound は指定された座席が存在しない場合に返される。
var ErrSeatNotFound = errors.New("seat not found")

// UseCase handles QR code business logic
type UseCase struct {
	db *sql.DB
}

// NewUseCase creates a new QR code use case
func NewUseCase(db *sql.DB) *UseCase {
	return &UseCase{db: db}
}

// GetQRCode processes the QR code request and logs the seat number
func (uc *UseCase) GetQRCode(req *qrcode.QRCodeRequest) (string, error) {
	// seat_idを整数に変換
	seatID, err := strconv.Atoi(req.SeatNumber)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidSeatNumber, req.SeatNumber)
	}

	// seat_idを条件にseatsテーブルをSELECT
	var id int
	var createdAt, updatedAt string
	query := "SELECT id, created_at, updated_at FROM seats WHERE id = ?"
	err = uc.db.QueryRow(query, seatID).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Seat not found: seat_id=%d", seatID)
			return "", fmt.Errorf("%w: seat_id=%d", ErrSeatNotFound, seatID)
		}
		return "", fmt.Errorf("failed to query seat: %w", err)
	}

	// 結果をログ出力
	log.Printf("Seat found - ID: %d, CreatedAt: %s, UpdatedAt: %s", id, createdAt, updatedAt)

	return fmt.Sprintf("Seat found - ID: %d, CreatedAt: %s, UpdatedAt: %s", id, createdAt, updatedAt), nil
}
