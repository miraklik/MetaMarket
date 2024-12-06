package handlers

import (
	"net/http"
	"nft-marketplace/services"
	"nft-marketplace/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
func GetNFTs(db *gorm.DB, ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Raw("SELECT * FROM nfts").Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		defer rows.Close()

		var nfts []map[string]interface{}
		for rows.Next() {
			var id int
			var tokenID string
			var ownerAddress string
			if err := rows.Scan(&id, &tokenID, &ownerAddress); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan rows"})
				return
			}
			nfts = append(nfts, map[string]interface{}{
				"id":            id,
				"token_id":      tokenID,
				"owner_address": ownerAddress,
			})
		}

		c.JSON(http.StatusOK, gin.H{"data": nfts})
	}
}

// MintNFT handles the minting of a new NFT.
// It expects a JSON request with a recipient Ethereum address and a token ID.
// The function validates the recipient address and uses the EthereumService to mint the NFT.
// If successful, it inserts the new NFT record into the database.
// Responds with a success message if the operation is completed without errors,
// otherwise, it provides an appropriate error response.
func MintNFT(db *gorm.DB, ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Recipient string `json:"recipient"`
			TokenID   string `json:"token_id"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := utils.ValidateEthereumAddress(request.Recipient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipient address"})
			return
		}

		if err := ethService.MintNFT(request.Recipient); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		query := "INSERT INTO nfts (token_id, owner_address) VALUES ($1, $2)"
		if err := db.Exec(query, request.TokenID, request.Recipient); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "NFT minted successfully"})
	}
}

// BuyNFT handles the purchase of an existing NFT.
// It expects a JSON request with the token ID of the NFT and the Ethereum address of the buyer.
// The function validates the buyer address and uses the EthereumService to transfer the NFT to the buyer.
// If successful, it updates the owner address of the NFT in the database.
// Responds with a success message if the operation is completed without errors,
// otherwise, it provides an appropriate error response.
func BuyNFT(db *gorm.DB, ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request struct {
			TokenID string `json:"token_id"`
			Buyer   string `json:"buyer"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
			return
		}

		if err := utils.ValidateEthereumAddress(request.Buyer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid buyer address: " + err.Error()})
			return
		}

		var currentOwner string
		query := "SELECT owner_address FROM nfts WHERE token_id = $1"
		if err := db.Raw(query, request.TokenID).Scan(&currentOwner).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "NFT not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
			return
		}

		if err := ethService.TransferNFT(currentOwner, request.Buyer, request.TokenID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to transfer NFT: " + err.Error()})
			return
		}

		updateQuery := "UPDATE nfts SET owner_address = $1 WHERE token_id = $2"
		if err := db.Exec(updateQuery, request.Buyer, request.TokenID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update database: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "NFT purchased successfully!"})
	}
}
