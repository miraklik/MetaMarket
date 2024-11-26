package middleware

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	AuthorizationHeader    = "Authorization"
	BearerPrefix           = "Bearer "
	UserRoleHeader         = "X-User-Role"
	BuyerIDHeader          = "X-Buyer-ID"
	ErrorUnauthorized      = "Unauthorized"
	ErrorInvalidToken      = "Invalid token"
	ErrorDatabase          = "Database error"
	ErrorNotFound          = "No NFTs found"
	ErrorInsufficientFunds = "Insufficient funds"
)

// parseToken parses the given token string and returns the underlying claims.
func parseToken(tokenStr, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New(ErrorInvalidToken)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

// getUserBalance gets the balance of a user from the database.
func getUserBalance(db *sql.DB, userID string) (float64, error) {
	var balance float64
	query := "SELECT balance FROM users WHERE id = $1"
	err := db.QueryRow(query, userID).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, sql.ErrNoRows
		}
		return 0, err
	}
	return balance, nil
}

// GetNFTs is a middleware function that checks the Authorization header for a valid Bearer token. If the token is valid, it queries the database for the NFTs associated with the user ID present in the token. If the query succeeds, it returns the list of NFTs in the response body. If the token is invalid or the query fails, it returns an appropriate error response.
func GetNFTs(secretKey string, db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader(AuthorizationHeader)
		if tokenStr == "" || !strings.HasPrefix(tokenStr, BearerPrefix) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrorUnauthorized})
			c.Abort()
			return
		}

		claims, err := parseToken(strings.TrimPrefix(tokenStr, BearerPrefix), secretKey)
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
		rows, err := db.Query(query, userID)
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
func BuyNFT(db *sql.DB) gin.HandlerFunc {
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
