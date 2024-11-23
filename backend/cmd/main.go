package main

import (
	"fmt"
	"log"
	"math/big"
	"nft-marketplace/config"
	"nft-marketplace/db"
	"nft-marketplace/handlers"
	"nft-marketplace/middleware"
	"nft-marketplace/services"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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

	client, err := ethclient.Dial(config.BlockChainRPC)
	if err != nil {
		log.Fatal("Failed to connect to Ethereum client:", err)
	}

	auth, err := bind.NewTransactorWithChainID(config.PrivateKey, nil, big.NewInt(1))
	if err != nil {
		log.Fatal("Failed to create transactor:", err)
	}

	contractAddress := common.HexToAddress(config.ContractAddress)
	contractInstance, err := services.NewNFT(contractAddress, client)
	if err != nil {
		log.Fatal("Failed to create contract instance:", err)
	}

	escrowID, err := contractInstance.InitiatePayment(auth, big.NewInt(100))
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Fprintf(w, "Payment initiated successfully with Escrow ID: %d", escrowID)

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
	r.GET("/nfts", middleware.GetNFTs(), handlers.GetNFTs(dbConn, ethereumService))
	r.POST("/nfts/mint", middleware.MintNFT(), handlers.MintNFT(dbConn, ethereumService))
	r.POST("/nfts/buy", middleware.BuyNFT(), handlers.BuyNFT(dbConn, ethereumService))

	log.Fatal(r.Run(":8080"))
}
