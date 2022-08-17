package handler

import (
	"github.com/Porthorian/urling.io/shortener"
	"github.com/Porthorian/urling.io/store"
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
	"fmt"
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

	newUrl := creationRequest.LongUrl
	if !strings.Contains("https://", creationRequest.LongUrl) && !strings.Contains("http://", creationRequest.LongUrl) {
		newUrl = fmt.Sprintf("http://%s", creationRequest.LongUrl)
	}

	shortUrl := shortener.GenerateShortLink(newUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	proto := c.Request.Header.Get("X-Forwarded-Proto")
	if proto != "" {
		scheme = proto
	}
	host := fmt.Sprintf("%s://%s", scheme, c.Request.Host)
	c.JSON(200, gin.H{
		"message": "shortened",
		"short_url": fmt.Sprintf("%s/%s", host, shortUrl),
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
	c.String(http.StatusOK, "pong")
}