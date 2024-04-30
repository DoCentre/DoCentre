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

type docDto struct {
	ID        uint   `json:"id" example:"1"`
	AuthorID  uint   `json:"author_id" example:"1"`
	Title     string `json:"title" example:"Important Document"`
	Status    string `json:"status" example:"EDIT"`
	CreatedAt string `json:"created_at" example:"2021-08-01T00:00:00Z"`
	UpdatedAt string `json:"updated_at" example:"2021-08-01T00:00:00Z"`
}

// @Summary Get author documents
// @Description Get all documents that belong to the author; the author has to be a existing user.
// @Tags Document
// @Accept json
// @Produce json
// @Param body body controllers.GetAuthorDocuments.requestBody true " "
// @Success 200 {object} controllers.GetAuthorDocuments.successResponseBody
// @Failure 400 {object} controllers.GetAuthorDocuments.invalidResponseBody
// @Failure 500 {object} controllers.GetAuthorDocuments.failedResponseBody
// @Router /documents/author [post]
func GetAuthorDocuments(c *gin.Context) {
	type requestBody struct {
		AuthorID uint `json:"author_id" example:"1"`
	}
	type invalidResponseBody struct {
		Error string `json:"error" example:"Invalid request body"`
	}
	type failedResponseBody struct {
		Error string `json:"error" example:"Failed to get documents"`
	}
	type successResponseBody struct {
		Documents []docDto `json:"documents"`
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

	docs, err := services.GetAuthorDocuments(body.AuthorID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, failedResponseBody{
			Error: "Failed to get documents",
		})
		return
	}

	var docsDto []docDto

	for _, doc := range docs {
		docsDto = append(docsDto, docDto{
			ID:        doc.ID,
			AuthorID:  doc.AuthorID,
			Title:     doc.Title,
			Status:    doc.Status,
			CreatedAt: doc.CreatedAt.String(),
			UpdatedAt: doc.UpdatedAt.String(),
		})
	}

	c.JSON(http.StatusOK, successResponseBody{
		Documents: docsDto,
	})
}

// @Summary Get viewer documents
// @Description Get all documents that belong to the viewer; the viewer has to be a existing user.
// @Tags Document
// @Accept json
// @Produce json
// @Param body body controllers.GetViewerDocuments.requestBody true " "
// @Success 200 {object} controllers.GetViewerDocuments.successResponseBody
// @Failure 400 {object} controllers.GetViewerDocuments.invalidResponseBody
// @Failure 500 {object} controllers.GetViewerDocuments.failedResponseBody
// @Router /documents/viewer [post]
func GetViewerDocuments(c *gin.Context) {
	type requestBody struct {
		ViewerID uint `json:"viewer_id" example:"1"`
	}
	type invalidResponseBody struct {
		Error string `json:"error" example:"Invalid request body"`
	}
	type failedResponseBody struct {
		Error string `json:"error" example:"Failed to get documents"`
	}
	type successResponseBody struct {
		Documents []docDto `json:"documents"`
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

	docs, err := services.GetViewerDocuments(body.ViewerID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, failedResponseBody{
			Error: "Failed to get documents",
		})
		return
	}

	var docsDto []docDto

	for _, doc := range docs {
		docsDto = append(docsDto, docDto{
			ID:        doc.ID,
			AuthorID:  doc.AuthorID,
			Title:     doc.Title,
			Status:    doc.Status,
			CreatedAt: doc.CreatedAt.String(),
			UpdatedAt: doc.UpdatedAt.String(),
		})
	}

	c.JSON(http.StatusOK, successResponseBody{
		Documents: docsDto,
	})
}

// @Summary Add viewer
// @Description Add a viewer to the document; the viewer has to be a existing user.
// @Tags Document
// @Accept json
// @Produce json
// @Param body body controllers.AddViewer.requestBody true " "
// @Success 200 {object} controllers.AddViewer.successResponseBody
// @Failure 400 {object} controllers.AddViewer.invalidResponseBody
// @Failure 500 {object} controllers.AddViewer.failedResponseBody
// @Router /document/add/viewer [post]
func AddViewer(c *gin.Context) {
	type requestBody struct {
		DocumentID uint `json:"document_id" binding:"required" example:"1"`
		ViewerID   uint `json:"viewer_id" binding:"required" example:"1"`
	}
	type invalidResponseBody struct {
		Error string `json:"error" example:"Invalid request body"`
	}
	type failedResponseBody struct {
		Error string `json:"error" example:"Failed to add viewer to the document"`
	}
	type successResponseBody struct{}

	var body requestBody
	err := c.BindJSON(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, invalidResponseBody{
			Error: "Invalid request body",
		})
		return
	}

	err = services.AddViewer(body.DocumentID, body.ViewerID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, failedResponseBody{
			Error: "Failed to add viewer to the document",
		})
		return
	}

	c.JSON(http.StatusOK, successResponseBody{})
}
