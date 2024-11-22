package handlers

import (
	"database/sql"
	"net/http"
	"nft-marketplace/services"

	"github.com/gin-gonic/gin"
)

func GetNFTs(db *sql.DB, ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM nfts")
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

func MintNFT(db *sql.DB, ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Recipient string `json:"recipient"`
			TokenID   string `json:"token_id"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := services.ValidateEthereumAddress(request.Recipient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipient address"})
			return
		}

		if err := ethService.MintNFT(request.Recipient); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		query := "INSERT INTO nfts (token_id, owner_address) VALUES ($1, $2)"
		_, err := db.Exec(query, request.TokenID, request.Recipient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "NFT minted successfully"})
	}
}

func BuyNFT(db *sql.DB, ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request struct {
			TokenID string `json:"token_id"`
			Buyer   string `json:"buyer"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
			return
		}

		if err := services.ValidateEthereumAddress(request.Buyer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid buyer address: " + err.Error()})
			return
		}

		var currentOwner string
		query := "SELECT owner_address FROM nfts WHERE token_id = $1"
		err := db.QueryRow(query, request.TokenID).Scan(&currentOwner)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "NFT not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
			return
		}

		if err := ethService.TransferNFT(currentOwner, request.Buyer, request.TokenID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to transfer NFT: " + err.Error()})
			return
		}

		updateQuery := "UPDATE nfts SET owner_address = $1 WHERE token_id = $2"
		_, err = db.Exec(updateQuery, request.Buyer, request.TokenID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update database: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "NFT purchased successfully!"})
	}
}
