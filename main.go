package main

import (
	"net/http"

	"github.com/docentre/docentre/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginswagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type healthResponse struct {
	Message string `json:"message"`
}

// @Summary Check health
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} healthResponse
// @Router /health [get]
func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, healthResponse{
		Message: "health check success",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.Title = "Docentre API"
	docs.SwaggerInfo.Description = "This is the API documentation for Docentre."
	r.GET("/health", CheckHealth)
	r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}
