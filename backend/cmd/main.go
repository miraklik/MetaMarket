package main

import (
	"log"
	"nft-marketplace/config"
	"nft-marketplace/db"
	"nft-marketplace/handlers"
	"nft-marketplace/services"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	dbConn, err := db.ConnectDB(config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer dbConn.Close()

	ethereumService, err := services.NewEthereumService(
		config.BlockChainRPC,
		config.ContractAddress,
		os.Getenv("PRIVATE_KEY"),
		config.MarketplaceABI,
	)
	if err != nil {
		log.Fatalf("Не удалось подключиться к контракту: %v", err)
	}

	r := gin.Default()
	r.GET("/nfts", handlers.GetNFTs(dbConn, ethereumService))
	r.POST("/nfts/mint", handlers.MintNFT(dbConn, ethereumService))
	r.POST("/nfts/buy", handlers.BuyNFT(dbConn, ethereumService))

	log.Fatal(r.Run(":8080"))
}
