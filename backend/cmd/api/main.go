package main

import (
	"log"
	"nft-marketplace/auth"
	"nft-marketplace/db"
	"nft-marketplace/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	db, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return db
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := DBInit()

	server := auth.NewServer(db)

	router := r.Group("/api")

	router.POST("/register", server.Register)
	router.POST("/login", server.Login)

	authorized := r.Group("/api/admin")

	authorized.Use(middleware.JwtAuthMiddleware())
	authorized.GET("/groceries", server.GetGroceries)
	authorized.POST("/groceries", server.PostGrocery)

	return r
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	port := os.Getenv("SERVER_ADDRESS")

	r := SetupRouter()

	log.Fatal(r.Run(":" + port))

}
