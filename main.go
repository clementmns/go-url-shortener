package main

import (
	"fmt"
	"os"
	"url-shortener/handler"
	"url-shortener/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT must be set")
	}

	r := gin.Default()
	// GET: Home
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "URL Shortener",
		})
	})

	// POST: Create a short URL
	r.POST("/url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// GET: Redirect to the original URL
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":" + port)
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %s", err.Error()))
	}
}
