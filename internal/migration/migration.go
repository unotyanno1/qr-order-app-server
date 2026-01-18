package migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

)

// RunMigrations はデータベースマイグレーションを実行します
func RunMigrations() error {
	// 環境変数からDB接続情報を取得
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "qr_order_db")

	// DSN (Data Source Name) を構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// データベース接続（リトライロジック付き）
	var db *sql.DB
	var err error
	maxRetries := 10
	retryInterval := 3 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			return fmt.Errorf("failed to open database: %w", err)
		}

		// データベース接続を確認
		if err = db.Ping(); err == nil {
			break
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(retryInterval)
		}
		db.Close()
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
	}
	defer db.Close()

	// MySQLドライバーインスタンスを作成
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to create mysql driver: %w", err)
	}

	// マイグレーションファイルのパス
	migrationsPath := getEnv("MIGRATIONS_PATH", "file://migrations")

	// migrateインスタンスを作成
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"mysql",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// マイグレーションを実行
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migration changes detected")
			return nil
		}
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

// getEnv は環境変数を取得し、存在しない場合はデフォルト値を返します
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
