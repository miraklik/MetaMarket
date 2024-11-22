package main

import (
	"log"
	"nft-marketplace/config"
	"nft-marketplace/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	dbConn, err := db.ConnectDB(config.)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer dbConn.Close()

	// Подключение к блокчейну
	blockchainClient := blockchain.NewClient(config.BlockChainRPC)
	marketplace := blockchain.NewMarketplace(blockchainClient, config.ContractAddress)

	if marketplace == nil {
		log.Fatalf("Не удалось подключиться к контракту")
	}

	r := gin.Default()
	r.POST("/list", func(c *gin.Context) {
		type ListNFTRequest struct {
			TokenID       int64   `json:"token_id" binding:"required"`
			SellerAddress string  `json:"seller_address" binding:"required"`
			Price         float64 `json:"price" binding:"required"`
		}

		var req ListNFTRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Price <= 0 {
			c.JSON(400, gin.H{"error": "Price must be greater than zero"})
			return
		}

		txHash, err := blockchain.ListNFTOnBlockchain(req.TokenID, req.SellerAddress, req.Price)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to list NFT on blockchain", "details": err.Error()})
			return
		}

		arg := db.CreateListingParams{
			NftID:         req.TokenID,
			SellerAddress: req.SellerAddress,
			Price:         req.Price,
			Status:        "active",
		}

		listing, err := db.CreateListing(c, arg)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save listing in database", "details": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "NFT listed successfully",
			"tx_hash": txHash,
			"listing": listing,
		})
	})

	log.Fatal(r.Run(":8080"))
}
