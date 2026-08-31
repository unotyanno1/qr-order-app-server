package database

import (
    "database/sql"
    "os"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func NewMySQL() (*sql.DB, error) {
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		dbName,
	)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("faild to open database: %w", err)
    }

    if err := db.Ping(); err != nil {
        db.Close()
        return nil, fmt.Errorf("failed to connect database: %w", err)
    }

    return db, nil
}