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
		DocumentID uint `json:"document_id" example:"1"`
	}

	var body requestBody
	err := c.BindJSON(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, invalidResponseBody{
			Error: "Invalid request body",
		})
		return
	}

	doc, err := services.CreateDocument(body.AuthorID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, failedResponseBody{
			Error: "Failed to create document",
		})
		return
	}

	c.JSON(http.StatusOK, successResponseBody{
		DocumentID: doc.ID,
	})
}

// @Summary Update document
// @Description Update a document that belongs to the author; the author has to be a existing user. if no approver yet, approver_id should be 0.
// @Tags Document
// @Accept json
// @Produce json
// @Param body body controllers.UpdateDocument.requestBody true " "
// @Success 200 {object} controllers.UpdateDocument.successResponseBody
// @Failure 400 {object} controllers.UpdateDocument.invalidResponseBody
// @Failure 500 {object} controllers.UpdateDocument.failedResponseBody
// @Router /document/update [put]
func UpdateDocument(c *gin.Context) {
	type requestBody struct {
		DocumentID uint   `json:"document_id" binding:"required" example:"1"`
		AuthorID   uint   `json:"author_id" binding:"required" example:"1"`
		Title      string `json:"title" example:"Hello"`
		Content    string `json:"content" binding:"required" example:"Hello, world!"`
		Appendix   string `json:"appendix" example:""`
		Status     string `json:"status" example:"EDIT"`
		ApproverID uint   `json:"approver_id" example:"0"`
	}
	type invalidResponseBody struct {
		Error string `json:"error" example:"Invalid request body"`
	}
	type failedResponseBody struct {
		Error string `json:"error" example:"Failed to update document"`
	}
	type successResponseBody struct {
		DocumentID uint `json:"document_id" example:"1"`
	}

	var body requestBody
	err := c.BindJSON(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, invalidResponseBody{
			Error: "Invalid request body",
		})
		return
	}

	documentID, err := services.UpdateDocument(services.UpdateDocumentSnapshot{
		DocumentID: body.DocumentID,
		AuthorID:   body.AuthorID,
		Title:      body.Title,
		Content:    body.Content,
		Appendix:   body.Appendix,
		Status:     body.Status,
		ApproverID: body.ApproverID,
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, failedResponseBody{
			Error: "Failed to update document",
		})
		return
	}

	c.JSON(http.StatusOK, successResponseBody{
		DocumentID: documentID,
	})
}
