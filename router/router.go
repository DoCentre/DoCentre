package router

import (
	"net/http"

	"github.com/docentre/docentre/controllers"
	"github.com/docentre/docentre/docs"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginswagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @Summary Check health
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} router.CheckHealth.responseBody
// @Router /health [get]
func CheckHealth(c *gin.Context) {
	type responseBody struct {
		Message string `json:"message"`
	}

	c.JSON(http.StatusOK, responseBody{
		Message: "health check success",
	})
}

func setSwagger(r *gin.Engine) {
	docs.SwaggerInfo.Title = "Docentre API"
	docs.SwaggerInfo.Description = "This is the API documentation for Docentre."
	r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler, ginswagger.DefaultModelsExpandDepth(-1)))
}

func setHealthCheckController(r *gin.Engine) {
	r.GET("/health", CheckHealth)
}

func setUserController(r *gin.Engine) {
	r.POST("/user", controllers.UserCreate)
	r.POST("/login", controllers.UserLogin)
	r.POST("/users", controllers.GetUsersByUsername)
}

func setDocumentController(r *gin.Engine) {
	r.POST("/document", controllers.CreateDocument)
	// TODO r.PATCH("/documents/:document_id", controllers.UpdateDocument)
	// TODO r.PATCH("/documents/:document_id/status", controllers.SetDocumentStatus)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	setHealthCheckController(r)
	setUserController(r)
	setDocumentController(r)
	setSwagger(r)

	return r
}
