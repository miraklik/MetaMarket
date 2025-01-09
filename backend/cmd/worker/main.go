package main

import (
	"log"
	"nft-marketplace/config"
	"nft-marketplace/db"
	"nft-marketplace/handlers"
	"nft-marketplace/middleware"
	"nft-marketplace/services"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return db
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file:", err)
	}
	cfg := config.LoadConfig()

	db := InitDB()

	client, err := ethclient.Dial(cfg.BlockChainRPC)
	if err != nil {
		log.Panicf("Failed to connect to Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	etherService := &services.EthereumService{
		Client:          client,
		ContractAddress: common.HexToAddress(cfg.ContractAddress),
		PrivateKey:      privateKey,
		Contract:        nil,
	}

	router := gin.Default()

	server := handlers.NewServers(db)

	middlewareNFTs := router.Group("/nfts")

	middlewareNFTs.Use(middleware.MintNFT(etherService))
	router.POST("/Create", server.MintNFT(etherService))
	middlewareNFTs.Use(middleware.GetNFTs(etherService))
	router.GET("/nfts/:id", handlers.GetNFTs(etherService))
	middlewareNFTs.Use(middleware.BuyNFT(etherService))
	router.POST("/Buy", handlers.BuyNFT(etherService))
	router.GET("/Search", handlers.SearchNFTs(etherService))
	router.DELETE("/nfts/:id", handlers.DeleteNFT(etherService))

	router.Run(os.Getenv("SERVER_ADDRESS"))
}
