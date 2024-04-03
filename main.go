package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "health check success",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", CheckHealth)
	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}
