package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sqlx.DB

func ConnectDB() {
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "5432")
	user := GetEnv("DB_USER", "postgres")
	password := GetEnv("DB_PASSWORD", "")
	dbname := GetEnv("DB_NAME", "postgres")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL!")
}
