package middleware

import (
	"math/big"
	"net/http"
	"nft-marketplace/services"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// GetNFTs is a middleware function that retrieves a list of NFTs owned by the
// user whose address is provided in the Authorization header. The function
// expects a JSON web token in the Authorization header, which is validated
// before retrieving the NFTs. If the token is invalid or missing, it returns an
// error response. If the user is not found or there is an error during the
// smart contract query, it returns an error response. If the user has no NFTs,
// it returns a not found error response. Otherwise, it returns a list of NFTs
// owned by the user.
func GetNFTs(ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, err := parseToken(strings.TrimPrefix(tokenStr, "Bearer "))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		userAddress, ok := claims["user_address"].(string)
		if !ok || !common.IsHexAddress(userAddress) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing user address"})
			c.Abort()
			return
		}

		nfts, err := ethService.GetNFTs(common.HexToAddress(userAddress))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch NFTs: " + err.Error()})
			return
		}

		if len(nfts) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No NFTs found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"nfts": nfts})
	}
}

func MintNFT(ethService *services.EthereumService) gin.HandlerFunc {
	// MintNFT is a middleware function that mints a new NFT. It retrieves the user's address from the User-Address header
	// and verifies its presence. If the user address is missing, it responds with an unauthorized error.
	return func(c *gin.Context) {
		userAddress := c.GetHeader("User-Address")
		if !common.IsHexAddress(userAddress) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user address"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// BuyNFT is a middleware function that validates the buyer's address and checks if the buyer has sufficient balance
// to purchase the NFT. It retrieves the user's address from the Buyer-Address header and verifies its presence.
// If the user address is missing or invalid, it responds with an unauthorized error.
// If the buyer's balance is insufficient, it responds with a 402 error.
// If there is an error during the balance query, it responds with an internal server error.
func BuyNFT(ethService *services.EthereumService) gin.HandlerFunc {
	return func(c *gin.Context) {
		buyerAddress := c.GetHeader("Buyer-Address")
		if !common.IsHexAddress(buyerAddress) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid buyer address"})
			c.Abort()
			return
		}

		balance, err := ethService.GetBalance(common.HexToAddress(buyerAddress))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch balance: " + err.Error()})
			return
		}

		if balance.Cmp(big.NewInt(0)) <= 0 {
			c.JSON(http.StatusPaymentRequired, gin.H{"error": "Insufficient funds"})
			return
		}

		c.Next()
	}
}
