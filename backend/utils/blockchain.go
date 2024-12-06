package utils

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectEthereum(repURL string) (*ethclient.Client, error) {

	client, err := ethclient.Dial(repURL)
	if err != nil {
		log.Fatal("Failed to connect to Ethereum client:", err)
		return nil, err
	}

	log.Println("Successfully connected to Ethereum client")
	return client, nil
}

func CheckTransaction(client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err == ethereum.NotFound {
			time.Sleep(2 * time.Second)
			continue
		} else if err != nil {
			return nil, err
		}
		return receipt, nil
	}
}
