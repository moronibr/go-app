package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "123456"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3307"),
		getEnv("DB_NAME", "orderingdb"),
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao pingar banco:", err)
	}

	log.Println("✅ Conectado ao banco de dados!")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
