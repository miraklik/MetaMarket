package utils

import (
	"errors"
	"fmt"
	"log"
	"nft-marketplace/db"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user db.User) (string, error) {
	tokenLifespanStr := os.Getenv("TOKEN_HOUR_LIFESPAN")
	if tokenLifespanStr == "" {
		return "", fmt.Errorf("TOKEN_HOUR_LIFESPAN is not set")
	}

	tokenLifespan, err := strconv.Atoi(tokenLifespanStr)
	if err != nil {
		log.Printf("Error converting TOKEN_HOUR_LIFESPAN: %v", err)
		return "", err
	}

	claims := jwt.MapClaims{
		"authorized": true,
		"id":         user.ID,
		"exp":        time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	apiSecret := os.Getenv("API_SECRET")
	if apiSecret == "" {
		return "", fmt.Errorf("API_SECRET is not set")
	}

	signedToken, err := token.SignedString([]byte(apiSecret))
	if err != nil {
		log.Printf("Error signing the token: %v", err)
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(c *gin.Context) error {
	token, err := GetToken(c)
	if err != nil {
		log.Printf("Could not get token: %v", err)
		return err
	}
	if token == nil {
		return errors.New("token is nil")
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func GetToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	if tokenString == "" {
		return nil, errors.New("token is missing from the request")
	}

	apiSecret := os.Getenv("API_SECRET")
	if apiSecret == "" {
		return nil, fmt.Errorf("API_SECRET is not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(apiSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	return token, nil
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}

	return ""
}

func CurrentUser(c *gin.Context) (db.User, error) {
	err := ValidateToken(c)
	if err != nil {
		return db.User{}, err
	}

	token, err := GetToken(c)
	if err != nil {
		return db.User{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return db.User{}, errors.New("invalid claims")
	}
	userId := uint(claims["id"].(float64))

	user, err := db.GetUserById(userId)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}
