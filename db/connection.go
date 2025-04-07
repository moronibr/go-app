package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	dsn := "user:password@tcp(localhost:3307)/orderingdb"
	var err error

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open DB: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}

	fmt.Println("✅ Connected to MySQL!")
	return nil
}
