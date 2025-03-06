package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"nft-marketplace/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
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
func parseToken(tokenStr string) (jwt.MapClaims, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("error loading .env file")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
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

// JwtAuthMiddleware is a middleware function that validates the JWT token in the
// Authorization header. If the token is invalid or missing, it responds with a 401
// Unauthorized response. Otherwise, it calls the next handler in the chain.
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.ValidateToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
