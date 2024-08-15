package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lokesh-katari/21341A0571/controllers"
)

// Allowed companies and categories

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	router := gin.Default()
	// Define the route for fetching products
	router.GET("/products", controllers.GetProductsHandler)
	router.GET("/products/:productId", controllers.GetProductByIDHandler)

	// Start the server
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
