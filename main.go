package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	qrcodehandler "github.com/unotyanno1/qr-order-app-server/handler/qrcode"
	"github.com/unotyanno1/qr-order-app-server/internal/migration"
	qrcodeusecase "github.com/unotyanno1/qr-order-app-server/usecase/qrcode"
)

func main() {
	fmt.Println("Docker Hello world!")

	// データベースマイグレーションを実行
	log.Println("Running database migrations...")
	if err := migration.RunMigrations(); err != nil {
		log.Printf("Migration error: %v", err)
		// マイグレーションエラーでもアプリケーションは起動を続ける
		// 必要に応じて log.Fatal(err) に変更してください
	}

	// アプリ起動時に DB 接続プールを一度だけ生成
	db, err := openDB()
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize use case and handler
	qrCodeUseCase := qrcodeusecase.NewUseCase(db)
	qrCodeHandler := qrcodehandler.NewHandler(qrCodeUseCase)

	// Register routes
	e.GET("/qr_code/:seat_number", qrCodeHandler.GetQRCode)

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := e.Start(port); err != nil {
		log.Fatal(err)
	}
}

// openDB は環境変数から接続情報を読み取り、DB 接続プールを生成する。
func openDB() (*sql.DB, error) {
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
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// getEnv は環境変数を取得し、存在しない場合はデフォルト値を返す。
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
