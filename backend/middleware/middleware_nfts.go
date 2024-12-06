package middleware

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// getUserBalance gets the balance of a user from the database.
func getUserBalance(db *gorm.DB, userID string) (float64, error) {
	var balance float64
	query := "SELECT balance FROM users WHERE id = ?"

	result := db.Raw(query, userID).Scan(&balance)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return 0, gorm.ErrRecordNotFound
		}
		return 0, result.Error
	}

	return balance, nil
}

// GetNFTs is a middleware function that retrieves a list of NFTs owned by the authenticated user.
// It requires a valid JWT in the Authorization header with the Bearer prefix. The function
// extracts the user ID from the token claims and queries the database for NFTs associated with
// the user ID. If the token is missing, invalid, or the user ID is not found, it responds with
// an unauthorized error. If the query is successful but no NFTs are found, it responds with a
// not found error. Otherwise, it returns the list of NFTs with a status code 200.
func GetNFTs(secretKey string, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader(AuthorizationHeader)
		if tokenStr == "" || !strings.HasPrefix(tokenStr, BearerPrefix) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrorUnauthorized})
			c.Abort()
			return
		}

		claims, err := parseToken(strings.TrimPrefix(tokenStr, BearerPrefix))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid user_id in token"})
			c.Abort()
			return
		}

		query := `SELECT nft_name FROM nfts WHERE owner_id = $1`
		rows, err := db.Raw(query, userID).Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ErrorDatabase})
			return
		}
		defer rows.Close()

		var nfts []string
		for rows.Next() {
			var nft string
			if err := rows.Scan(&nft); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read data"})
				return
			}
			nfts = append(nfts, nft)
		}

		if len(nfts) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrorNotFound})
			return
		}

		c.JSON(http.StatusOK, gin.H{"nfts": nfts})
	}
}

// MintNFT is a middleware function that checks if the user has an 'admin' role.
// If the user role is not 'admin', it responds with an unauthorized error and aborts the request.
// Otherwise, it allows the request to proceed to the next handler.
func MintNFT() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetHeader(UserRoleHeader)
		if userRole != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrorUnauthorized})
			c.Abort()
			return
		}
		c.Next()
	}
}

// BuyNFT is a middleware function that checks the buyer's ID and balance.
// It retrieves the buyer's ID from the request header and verifies its presence.
// If the buyer ID is missing, it responds with an unauthorized error.
// It also checks the buyer's balance from the database. If the buyer is not found or there's
// a database error, it responds with the appropriate error message.
// If the buyer's balance is zero or negative, it responds with a payment required error.
// If all validations pass, the request proceeds to the next handler.
func BuyNFT(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		buyerID := c.GetHeader(BuyerIDHeader)
		if buyerID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrorUnauthorized})
			c.Abort()
			return
		}

		buyerBalance, err := getUserBalance(db, buyerID)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Buyer not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": ErrorDatabase})
			return
		}

		if buyerBalance <= 0 {
			c.JSON(http.StatusPaymentRequired, gin.H{"error": ErrorInsufficientFunds})
			return
		}

		c.Next()
	}
}
