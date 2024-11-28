package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string `mapstructure:"DBHost"`
	DBName string `mapstructure:"DBName"`
	DBPort string `mapstructure:"DBPort"`
	DBUser string `mapstructure:"DBUser"`
	DBPass string `mapstructure:"DBPassword"`

	ServerAddress string `mapstructure:"SERVER_ADDRESS"`

	BlockChainRPC   string `mapstructure:"BLOCKCHAIN_RPC"`
	PrivateKey      string `mapstructure:"PRIVATE_KEY"`
	MarketplaceABI  string `mapstructure:"MARKETPLACE_ABI"`
	ContractAddress string `mapstructure:"CONTRACT_ADDRESS"`

	IPFSNodeAddress string `mapstructure:"IPFS_NODE_ADDRESS"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	return &Config{
		DBHost:          os.Getenv("DBHost"),
		DBName:          os.Getenv("DBName"),
		DBPort:          os.Getenv("DBPort"),
		DBUser:          os.Getenv("DBUser"),
		DBPass:          os.Getenv("DBPassword"),
		ServerAddress:   os.Getenv("SERVER_ADDRESS"),
		BlockChainRPC:   os.Getenv("BLOCKCHAIN_RPC"),
		PrivateKey:      os.Getenv("PRIVATE_KEY"),
		MarketplaceABI:  os.Getenv("MARKETPLACE_ABI"),
		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
		IPFSNodeAddress: os.Getenv("IPFS_NODE_ADDRESS"),
	}
}
