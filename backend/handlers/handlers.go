package handlers

import (
	"fmt"
	"log"
	"net/http"
	"nft-marketplace/db"
	"nft-marketplace/services"
	"nft-marketplace/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DB_Server struct {
	db *gorm.DB
}

func NewServers(db *gorm.DB) *DB_Server {
	return &DB_Server{db: db}
}

// GetNFTs returns a list of NFTs from the database in the following format:
// [
//
//	{
//	  "id": 1,
//	  "token_id": "1",
//	  "owner_address": "0x1234567890123456789012345678901234567890"
//	},
//	...
//
// ]
//
// The function handles database errors by returning an error response with status
// code 500. If the database query is successful, it returns the list of NFTs with
// status code 200.
func GetAllNFTs(ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Accounts string `json:"accounts"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if request.Accounts == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Account address is required"})
			return
		}

		if err := utils.ValidateEthereumAddress(request.Accounts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid owner address"})
			return
		}

		accounts := common.HexToAddress(request.Accounts)
		if ethService.ContractAddress == (common.Address{}) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Contract address not configured"})
			return
		}

		nfts, err := ethService.GetAllNFTs(accounts)
		if err != nil {
			log.Printf("Error fetching NFTs: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch NFTs: %v", err)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": nfts})
	}
}

// MintNFT is a handler function that mints a new NFT with the given token ID and recipient address.
// The function expects a JSON request with the following fields:
//
// - recipient: the Ethereum address of the recipient
// - token_id: the token ID of the NFT to be minted
//
// If the request is invalid or the recipient address is invalid, it responds with a bad request error.
// If there is an error during the smart contract call, it responds with an internal server error.
// If the database query fails, it responds with an internal server error.
// If the operation is successful, it responds with a success message with status code 200.
func (s *DB_Server) MintNFT(ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			TokenID     string `json:"token_id"`
			Name        string `json:"name"`
			Symbol      string `json:"symbol"`
			Description string `json:"description"`
			Price       string `json:"price"`
			Recipient   string `json:"recipient"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := utils.ValidateEthereumAddress(request.Recipient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipient address"})
			return
		}

		if err := utils.ValidatePrice(request.Price); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
			return
		}

		switch {
		case request.Name == "":
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
			return
		case request.Symbol == "":
			c.JSON(http.StatusBadRequest, gin.H{"error": "Symbol is required"})
			return
		case request.Description == "":
			c.JSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
			return
		}

		nfts := db.Nfts{
			Name:        request.Name,
			Symbol:      request.Symbol,
			Description: request.Description,
			Price:       request.Price,
		}

		err := ethService.MintNFT(request.TokenID, request.Price, request.Recipient)
		if err != nil {
			log.Printf("MintNFT error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mint NFT on blockchain: " + err.Error()})
			return
		}

		if err := s.db.Create(&nfts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create NFT: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "NFT minted successfully"})
	}
}

// BuyNFT handles the purchase of an NFT by transferring ownership from the current owner to the buyer.
// The function expects a JSON request containing the token ID of the NFT and the buyer's Ethereum address.
// It performs the following steps:
// 1. Validates the JSON request structure and the buyer's Ethereum address.
// 2. Retrieves the current owner of the NFT from the database.
// 3. Transfers the NFT using the EthereumService.
// 4. Updates the database with the new owner address.
// 5. Returns appropriate error responses if any step fails, including validation, database, or transfer errors.
// If successful, it responds with a status code 200 and a success message.func BuyNFT(ethService *services.EthereumService) gin.HandlerFunc {
func BuyNFT(ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			TokenID string `json:"token_id"`
			Buyer   string `json:"buyer"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := utils.ValidateEthereumAddress(request.Buyer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid buyer address"})
			return
		}

		err := ethService.TransferNFT(request.TokenID, request.Buyer)
		if err != nil {
			log.Printf("Error during NFT transfer: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to transfer NFT: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "NFT purchased successfully"})
	}
}

// Search is a handler function that searches for NFTs with the given name.
// The function expects a JSON request with a single field "name" containing the search query.
// It returns a list of NFTs with the given name, or an error if the search fails.
// The response is a JSON object with a single field "data" containing the list of NFTs.
// If the search is successful, it returns a status code 200.
// If the request is invalid or the search fails, it returns an appropriate error response.
func SearchNFTs(ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Name string `json:"name"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		nfts, err := ethService.SearchNFTs(request.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch NFTs: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": nfts})
	}
}

// DeleteNFT is a handler function that deletes an NFT with the given token ID.
// The function expects a JSON request with a single field "token_id" containing the token ID of the NFT to be deleted.
// It returns a success message with status code 200 if the deletion is successful.
// If the request is invalid or the deletion fails, it returns an appropriate error response.
func DeleteNFT(ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			TokenID string `json:"token_id"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		}

		if err := ethService.DeleteNFT(request.TokenID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete NFT: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "NFT deleted successfully"})
	}
}
