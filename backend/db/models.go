package db

import (
	"fmt"
	"log"
	"nft-marketplace/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	cfg := config.LoadConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}

	if err := db.AutoMigrate(&Nfts{}); err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}

	log.Println("Connected to the database")
	return db, nil
}
