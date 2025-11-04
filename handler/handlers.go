package handler

import (
	"net/http"
	"url-shortener/shortener"
	"url-shortener/store"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Url    string `json:"url" binding:"required"`
	UserId string `json:"userId" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	shortUrl := shortener.GenerateShortUrl(request.Url, request.UserId)
	store.SaveUrlMapping(shortUrl, request.Url)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message": "Short URL created",
		"url":     host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	url := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(http.StatusMovedPermanently, url)
}
