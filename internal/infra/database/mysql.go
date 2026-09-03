package database

import (
	"context"
	"database/sql"
	"fmt"

	"time"

	"github.com/unotyanno1/qr-order-app-server/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL(cfg config.Config) (*sql.DB, error) {
	dsn := buildDSN(cfg)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return db, nil
}

func buildDSN(cfg config.Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&timeout=5s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
}
