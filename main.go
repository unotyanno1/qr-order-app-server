package main

import (
	"fmt"
	"log"

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

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize use case and handler
	qrCodeUseCase := qrcodeusecase.NewUseCase()
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