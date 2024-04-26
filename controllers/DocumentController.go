package controllers

import (
	"log"
	"net/http"

	"github.com/docentre/docentre/services"
	"github.com/gin-gonic/gin"
)

// @Summary Create document
// @Description Create a new document that belongs to the author; the author has to be a existing user.
// @Tags Document
// @Accept json
// @Produce json
// @Param body body controllers.CreateDocument.requestBody true " "
// @Success 200 {object} controllers.CreateDocument.successResponseBody
// @Failure 400 {object} controllers.CreateDocument.invalidResponseBody
// @Failure 500 {object} controllers.CreateDocument.failedResponseBody
// @Router /document [post]
func CreateDocument(c *gin.Context) {
	type requestBody struct {
		AuthorID uint `json:"author_id" binding:"required" example:"1"`
	}
	type invalidResponseBody struct {
		Error string `json:"error" example:"Invalid request body"`
	}
	type failedResponseBody struct {
		Error string `json:"error" example:"Failed to create document"`
	}
	type successResponseBody struct {
		DocumentID uint `json:"document_id" example:"10"`
	}

	var body requestBody
	err := c.BindJSON(&body)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusBadRequest, invalidResponseBody{
			Error: "Invalid request body",
		})
		return
	}

	doc, err := services.CreateDocument(body.AuthorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, failedResponseBody{
			Error: "Failed to create document",
		})
		return
	}

	c.JSON(http.StatusOK, successResponseBody{
		DocumentID: doc.ID,
	})
}
