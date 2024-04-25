package controllers

import (
	"log"
	"net/http"

	"github.com/docentre/docentre/services"
	"github.com/gin-gonic/gin"
)

func CreateDocument(c *gin.Context) {
	var body struct {
		AuthorID uint `json:"author_id" binding:"required"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	doc, err := services.CreateDocument(body.AuthorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create document",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"document_id": doc.ID,
	})
}
