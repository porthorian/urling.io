package handler

import (
	"github.com/Porthorian/urling.io/shortener"
	"github.com/Porthorian/urling.io/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	host := "http://127.0.0.1:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host +shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	if initialUrl == nil {
		c.AbortWithStatus(404)
		return
	}

	c.Redirect(302, *initialUrl)
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "ping")
}