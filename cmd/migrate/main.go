package main

import (
	"log"
	"os"

	"github.com/unotyanno1/qr-order-app-server/internal/migration"
)

func main() {
	log.Println("Running database migrations...")
	if err := migration.RunMigrations(); err != nil {
		log.Printf("Migration error: %v", err)
		os.Exit(1)
	}
	log.Println("Migrations completed successfully")
}
