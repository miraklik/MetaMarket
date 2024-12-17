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
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	etherService := &services.EthereumService{
		Client:          client,
		ContractAddress: common.HexToAddress(cfg.ContractAddress),
		PrivateKey:      privateKey,
		Contract:        nil,
	}

	ctx, cancel := context.WithCancel(context.Background())
	go processMarketplaceOperations(ctx, client, *cfg)

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
			createListing(marketplaceInstance, client, cfg, big.NewInt(1), big.NewInt(1000000000000000000), "Test NFT", "This is a test NFT", "TEST")
			purchaseListing(marketplaceInstance, client, cfg, big.NewInt(1))
			time.Sleep(time.Minute)
		}
	}
}

func createListing(marketplaceInstance *marketplace.Marketplace, client *ethclient.Client, cfg config.Config, tokenID *big.Int, price *big.Int, name string, description string, symbol string) {
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

	tx, err := marketplaceInstance.CreateListing(auth, tokenID, price, name, description, symbol)
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

func purchaseListing(marketplaceInstance *marketplace.Marketplace, client *ethclient.Client, cfg config.Config, listingID *big.Int) {
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
