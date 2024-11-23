package middleware

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNFTs() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || token != "Bearer <token>" { // TODO: Добавить свою логику проверки
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func MintNFT() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetHeader("X-User-Role")
		if userRole != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func BuyNFT(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		buyerID := c.GetHeader("X-Buyer-ID")
		if buyerID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		var buyerBalance float64
		query := "SELECT balance FROM users WHERE id = $1"
		err := db.QueryRow(query, buyerID).Scan(&buyerBalance)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Buyer not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch buyer balance"})
			return
		}

		if buyerBalance <= 0 {
			c.JSON(http.StatusPaymentRequired, gin.H{"error": "Insufficient funds"})
			return
		}

		c.Next()
	}
}
