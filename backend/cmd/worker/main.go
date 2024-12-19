package main

import (
	"log"
	"nft-marketplace/config"
	"nft-marketplace/handlers"
	"nft-marketplace/middleware"
	"nft-marketplace/services"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file:", err)
	}
	cfg := config.LoadConfig()

	client, err := ethclient.Dial(cfg.BlockChainRPC)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
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

	middlewareNFTs := router.Group("/nfts")

	middlewareNFTs.Use(middleware.MintNFT(etherService))
	router.POST("/Create", handlers.MintNFT(etherService))
	middlewareNFTs.Use(middleware.GetNFTs(etherService))
	router.GET("/nfts/:id", handlers.GetNFTs(etherService))
	middlewareNFTs.Use(middleware.BuyNFT(etherService))
	router.POST("/Buy", handlers.BuyNFT(etherService))
	router.GET("/Search", handlers.SearchNFTs(etherService))

	router.Run(os.Getenv("SERVER_ADDRESS"))
}
