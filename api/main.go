package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// config.LoadConfig()
	// config.ConnectDatabase()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running"})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
