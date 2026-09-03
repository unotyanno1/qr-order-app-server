package database

import (
	"testing"

	"github.com/unotyanno1/qr-order-app-server/internal/config"
)

func TestBuildDSN(t *testing.T) {
	cfg := config.Config{
		DBUser: "root",
		DBPassword: "password",
		DBHost: "db",
		DBPort: "3306",
		DBName: "qr_order_db",
	}

	dsn := buildDSN(cfg)
	
	expected := "root:password@tcp(db:3306)/qr_order_db?parseTime=true&timeout=5s"
	if dsn != expected {
		t.Errorf("expected %s, got %s", expected, dsn)
	}
}