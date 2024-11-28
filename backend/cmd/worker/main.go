package main

import (
	"context"
	"log"
	"math/big"
	"nft-marketplace/accounts"
	marketplace "nft-marketplace/blockchain"
	"nft-marketplace/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	cfg := config.LoadConfig()

	client, err := ethclient.Dial(cfg.BlockChainRPC)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress(cfg.ContractAddress)

	marketplaceInstance, err := marketplace.NewMarketplace(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	listingCount, err := marketplaceInstance.ListingCount(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		log.Fatalf("Failed to get listing count: %v", err)
	}
	log.Printf("Listing count: %d", listingCount)

	title := "Test NFT"
	description := "Test NFT Description"
	imageHash := "QmTestIPFSHash"
	price := big.NewInt(1000000000000000000)

	auth, err := accounts.GetTransactor(cfg.PrivateKey, client)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	tx, err := marketplaceInstance.CreateListing(auth, title, description, imageHash, price)
	if err != nil {
		log.Fatalf("Failed to create listing: %v", err)
	}
	log.Printf("Listing created. Transaction hash: %s", tx.Hash().Hex())

	listingID := big.NewInt(1)
	tx, err = marketplaceInstance.PurchaseListing(auth, listingID)
	if err != nil {
		log.Fatalf("Failed to purchase listing: %v", err)
	}
	log.Printf("Purchase completed. Transaction hash: %s", tx.Hash().Hex())
}
