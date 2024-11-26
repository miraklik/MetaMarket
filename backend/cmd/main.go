package main

import (
	"io"
	"log"
	"math/big"
	"nft-marketplace/config"
	"nft-marketplace/db"
	"nft-marketplace/handlers"
	"nft-marketplace/middleware"
	"nft-marketplace/services"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	privateKeyStr := config.PrivateKey

	file, err := os.Open("./MarketplaceABI.json")
	if err != nil {
		log.Fatalf("Failed to open ABI file: %v", err)
	}
	defer file.Close()

	abiBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
	}

	contractABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	dbConn, err := db.ConnectDB(config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	client, err := ethclient.Dial(config.BlockChainRPC)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	chainID := big.NewInt(1)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	contractAddress := common.HexToAddress(config.ContractAddress)
	contractInstance := bind.NewBoundContract(contractAddress, contractABI, client, client, client)

	tx, err := contractInstance.Transact(auth, "purchaseListing", big.NewInt(100))
	if err != nil {
		log.Fatalf("Failed to initiate payment: %v", err)
		return
	}

	log.Printf("Payment initiated successfully. Transaction Hash: %s", tx.Hash().Hex())

	ethereumService, err := services.NewEthereumService(
		config.BlockChainRPC,
		config.ContractAddress,
		os.Getenv("PRIVATE_KEY"),
		config.MarketplaceABI,
	)
	if err != nil {
		log.Fatalf("Failed to create Ethereum service: %v", err)
	}

	r := gin.Default()
	r.GET("/nfts", middleware.GetNFTs(privateKeyStr, dbConn), handlers.GetNFTs(dbConn, ethereumService))
	r.POST("/nfts/mint", middleware.MintNFT(), handlers.MintNFT(dbConn, ethereumService))
	r.POST("/nfts/buy", middleware.BuyNFT(dbConn), handlers.BuyNFT(dbConn, ethereumService))

	log.Fatal(r.Run(":8080"))
}
