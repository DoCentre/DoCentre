package controllers

import (
	"log"
	"net/http"

	"github.com/docentre/docentre/services"
	"github.com/gin-gonic/gin"
)

type createDocumentRequest struct {
	AuthorID uint `json:"author_id" binding:"required" example:"1"`
}
type createDocumentInvalidResponse struct {
	Error string `json:"error" example:"Invalid request body"`
}
type createDocumentFailedResponse struct {
	Error string `json:"error" example:"Failed to create document"`
}
type createDocumentSuccessResponse struct {
	DocumentID uint `json:"document_id" example:"10"`
}

// @Summary Create document
// @Description Create a new document that belongs to the author; the author has to be a existing user.
// @Tags document
// @Accept json
// @Produce json
// @Param body body createDocumentRequest true " "
// @Success 200 {object} createDocumentSuccessResponse
// @Failure 400 {object} createDocumentInvalidResponse
// @Failure 500 {object} createDocumentFailedResponse
// @Router /document [post]
func CreateDocument(c *gin.Context) {
	var body createDocumentRequest
	err := c.BindJSON(&body)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusBadRequest, createDocumentInvalidResponse{
			Error: "Invalid request body",
		})
		return
	}

	doc, err := services.CreateDocument(body.AuthorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, createDocumentFailedResponse{
			Error: "Failed to create document",
		})
		return
	}

	c.JSON(http.StatusOK, createDocumentSuccessResponse{
		DocumentID: doc.ID,
	})
}
