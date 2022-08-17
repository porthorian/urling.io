package main

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/Porthorian/urling.io/handler"
	"github.com/Porthorian/urling.io/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	socket := ":9808"

	r.Use(CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("Hello from %s", strings.Split(c.Request.Host, socket)[0]))
	})

	r.POST("/shorten", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	r.GET("/ping", func(c *gin.Context) {
		handler.Ping(c)
	})

	// Note that store initialization happens here
	store.InitializeStore()

	err := r.Run(socket)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}