package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		log.Println("Health check endpoint hit")
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return r
}
