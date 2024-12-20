package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string `mapstructure:"DB_HOST"`
	DBName string `mapstructure:"DB_NAME"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASSWORD"`

	ServerAddress string `mapstructure:"SERVER_ADDRESS"`

	BlockChainRPC   string `mapstructure:"BLOCKCHAIN_RPC"`
	PrivateKey      string `mapstructure:"PRIVATE_KEY"`
	MarketplaceABI  string `mapstructure:"MARKETPLACE_ABI"`
	ContractAddress string `mapstructure:"CONTRACT_ADDRESS"`

	IPFSNodeAddress string `mapstructure:"IPFS_NODE_ADDRESS"`

	TokenLifespan string `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	APISecret     string `mapstructure:"API_SECRET"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	return &Config{
		DBHost:          os.Getenv("DB_HOST"),
		DBName:          os.Getenv("DB_NAME"),
		DBPort:          os.Getenv("DB_PORT"),
		DBUser:          os.Getenv("DB_USER"),
		DBPass:          os.Getenv("DB_PASSWORD"),
		ServerAddress:   os.Getenv("SERVER_ADDRESS"),
		BlockChainRPC:   os.Getenv("BLOCKCHAIN_RPC"),
		PrivateKey:      os.Getenv("PRIVATE_KEY"),
		MarketplaceABI:  os.Getenv("MARKETPLACE_ABI"),
		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
		IPFSNodeAddress: os.Getenv("IPFS_NODE_ADDRESS"),
		TokenLifespan:   os.Getenv("TOKEN_HOUR_LIFESPAN"),
		APISecret:       os.Getenv("API_SECRET"),
	}
}
