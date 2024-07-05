package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/weather-api/config"
	"github.com/joshua468/weather-api/internal"
	"github.com/joshua468/weather-api/internal/middleware"
)

func main() {

	config.LoadEnv()

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/", internal.SayHello)

	// Run the application
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	log.Printf("Starting server on port %s", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
