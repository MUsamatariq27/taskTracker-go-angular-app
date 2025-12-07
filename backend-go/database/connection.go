package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() {

	er := godotenv.Load()

	if er != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	DB = db

	fmt.Printf("Connected to PostgreSQL successfully: %s", dbName)

}
