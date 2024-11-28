package utils

import (
	"log"

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
