package main

import (
	"context"
	"log"
	"math/big"
	"nft-marketplace/accounts"
	marketplace "nft-marketplace/blockchain"
	"nft-marketplace/config"
	"nft-marketplace/handlers"
	"nft-marketplace/middleware"
	"nft-marketplace/services"
	"nft-marketplace/utils"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

	ctx, cancel := context.WithCancel(context.Background())
	go processMarketplaceOperations(ctx, client, *cfg)

	router := gin.Default()

	middlewareNFTs := router.Group("/nfts")

	middlewareNFTs.Use(middleware.MintNFT(&services.EthereumService{}))
	router.POST("/Create", handlers.MintNFT(&services.EthereumService{}))
	middlewareNFTs.Use(middleware.GetNFTs(&services.EthereumService{}))
	router.POST("/collection/nfts", handlers.GetNFTs(&services.EthereumService{}))
	middlewareNFTs.Use(middleware.BuyNFT(&services.EthereumService{}))
	router.POST("/Buy", handlers.BuyNFT(&services.EthereumService{}))

	router.Run(os.Getenv("SERVER_ADDRESS"))
	cancel()
}

func processMarketplaceOperations(ctx context.Context, client *ethclient.Client, cfg config.Config) {
	contractAddress := common.HexToAddress(cfg.ContractAddress)
	marketplaceInstance, err := marketplace.NewMarketplace(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}
	for {
		select {
		case <-ctx.Done():
			log.Println("Marketplace operations stopped")
			return
		default:
			createListing(marketplaceInstance, client, cfg)
			purchaseListing(marketplaceInstance, client, cfg)
			time.Sleep(time.Minute)
		}
	}
}

func createListing(marketplaceInstance *marketplace.Marketplace, client *ethclient.Client, cfg config.Config) {
	auth, err := accounts.GetTransactor(cfg.PrivateKey, client)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	addressContract := common.HexToAddress(cfg.ContractAddress)

	gasPrice := big.NewInt(5000000)
	auth.GasPrice = gasPrice

	callMsg := ethereum.CallMsg{
		From: auth.From,
		To:   &addressContract,
		Data: []byte{},
	}
	estimatedGasLimit, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Fatalf("Failed to estimate gas: %v", err)
	}
	auth.GasLimit = uint64(float64(estimatedGasLimit) * 1.2)

	balance, err := client.BalanceAt(context.Background(), auth.From, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve balance: %v", err)
	}
	txCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(auth.GasLimit)))
	if balance.Cmp(txCost) < 0 {
		log.Fatalf("Insufficient funds: balance %s, required %s", balance, txCost)
	}

	title := "Test NFT"
	description := "TNFT"
	imageHash := "QmTestIPFSHash"
	price := big.NewInt(100000000)

	tx, err := marketplaceInstance.CreateListing(auth, title, description, imageHash, price)
	if err != nil {
		log.Fatalf("Failed to create listing: %v", err)
	}
	log.Printf("Listing created. Transaction hash: %s", tx.Hash().Hex())

	receipt, err := utils.CheckTransaction(client, tx)
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}
	if receipt.Status == 0 {
		log.Fatalf("Transaction failed. Hash: %s", tx.Hash().Hex())
	}
	log.Printf("Transaction succeeded. Hash: %s", tx.Hash().Hex())
}

func purchaseListing(marketplaceInstance *marketplace.Marketplace, client *ethclient.Client, cfg config.Config) {
	auth, err := accounts.GetTransactor(cfg.PrivateKey, client)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	auth.GasLimit = uint64(300000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice

	listingID := big.NewInt(1)
	tx, err := marketplaceInstance.PurchaseListing(auth, listingID)
	if err != nil {
		log.Fatalf("Failed to purchase listing: %v", err)
	}
	log.Printf("Purchase completed. Transaction hash: %s", tx.Hash().Hex())

	receipt, err := utils.CheckTransaction(client, tx)
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}
	if receipt.Status == 0 {
		log.Fatalf("Transaction failed. Hash: %s", tx.Hash().Hex())
	}
	log.Printf("Transaction succeeded. Hash: %s", tx.Hash().Hex())
}
