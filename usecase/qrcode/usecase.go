package qrcode

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/unotyanno1/qr-order-app-server/domain/qrcode"

)

// UseCase handles QR code business logic
type UseCase struct{}

// NewUseCase creates a new QR code use case
func NewUseCase() *UseCase {
	return &UseCase{}
}

// getDBConnection returns a database connection
func getDBConnection() (*sql.DB, error) {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "qr_order_db")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// getEnv returns environment variable or default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetQRCode processes the QR code request and logs the seat number
func (uc *UseCase) GetQRCode(req *qrcode.QRCodeRequest) (string, error) {
	// seat_idを整数に変換
	seatID, err := strconv.Atoi(req.SeatNumber)
	if err != nil {
		return "",fmt.Errorf("invalid seat number: %s", req.SeatNumber)
	}

	// データベース接続を取得
	db, err := getDBConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	// seat_idを条件にseatsテーブルをSELECT
	var id int
	var createdAt, updatedAt string
	query := "SELECT id, created_at, updated_at FROM seats WHERE id = ?"
	err = db.QueryRow(query, seatID).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Seat not found: seat_id=%d", seatID)
			return "", nil
		}
		return "", fmt.Errorf("failed to query seat: %w", err)
	}

	// 結果をログ出力
	log.Printf("Seat found - ID: %d, CreatedAt: %s, UpdatedAt: %s", id, createdAt, updatedAt)
	
	return fmt.Sprintf("Seat found - ID: %d, CreatedAt: %s, UpdatedAt: %s", id, createdAt, updatedAt), nil
}
