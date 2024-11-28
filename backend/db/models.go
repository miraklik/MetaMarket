package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(host, port, user, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
		return nil, err
	}

	log.Println("Connected to the database")
	return db, nil
}

func CloseDB(db *sql.DB) error {
	if db == nil {
		log.Println("Database connection is nil")
		return nil
	}

	err := db.Close()
	if err != nil {
		log.Println("Failed to close the database connection:", err)
		return err
	}

	log.Panicln("Database connection closed")
	return nil
}

func Setup() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}
	if err = db.AutoMigrate(&Grocery{}); err != nil {
		log.Println(err)
	}

	return db, err
}
