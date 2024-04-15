package router

import (
	"net/http"

	"github.com/docentre/docentre/controllers"
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

func setSwagger(r *gin.Engine) {
	docs.SwaggerInfo.Title = "Docentre API"
	docs.SwaggerInfo.Description = "This is the API documentation for Docentre."
	r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}

func setHealthCheckController(r *gin.Engine) {
	r.GET("/health", CheckHealth)
}

func setUserController(r *gin.Engine) {
	r.POST("/user", controllers.UserCreate)
	r.POST("/login", controllers.UserLogin)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	setHealthCheckController(r)
	setUserController(r)

	setSwagger(r)

	return r
}
