package main

import (
	"log"

	"github.com/unotyanno1/qr-order-app-server/internal/config"
	"github.com/unotyanno1/qr-order-app-server/internal/infra/database"
)

func main() {
	cfg := config.Load()

	db, err := database.NewMySQL(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("database connected")
}
