package db

import (
	"database/sql"
	"fmt"
	"log"
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
