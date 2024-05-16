package controllers

import (
	"log"
	"net/http"

	"github.com/docentre/docentre/models"
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

// @Summary Get approver documents
// @Description Get all documents that need to be approved by the approver; the approver has to be a existing user.
// @Tags Document
// @Accept json
// @Produce json
// @Param body body controllers.GetApproverDocuments.requestBody true " "
// @Success 200 {object} controllers.GetApproverDocuments.successResponseBody
// @Failure 400 {object} controllers.GetApproverDocuments.invalidResponseBody
// @Failure 500 {object} controllers.GetApproverDocuments.failedResponseBody
// @Router /documents/approver [post]
func GetApproverDocuments(c *gin.Context) {
	type requestBody struct {
		ApproverID uint `json:"approver_id" example:"1"`
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

	docs, err := services.GetApproverDocuments(body.ApproverID)
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

// @Summary Get verify documents
// @Description Get all documents that has been verified, thus can be viewed by any user; if the user is an admin, return all documents instead.
// @Tags Document
// @Accept json
// @Produce json
// @Param body body controllers.GetVerifyDocuments.requestBody true " "
// @Success 200 {object} controllers.GetVerifyDocuments.successResponseBody
// @Failure 400 {object} controllers.GetVerifyDocuments.invalidResponseBody
// @Failure 500 {object} controllers.GetVerifyDocuments.failedResponseBody
// @Router /documents/verify [post]
func GetVerifyDocuments(c *gin.Context) {
	type requestBody struct {
		UserID uint `json:"user_id" example:"1"`
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

	docs, err := services.GetVerifyDocuments(body.UserID)
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

// @Summary Set document status
// @Description Set the status of the document; the approver has to be a existing user.
// @Tags Document
// @Accept json
// @Produce json
// @Param body body controllers.SetDocumentStatus.requestBody true " "
// @Success 200 {object} controllers.SetDocumentStatus.successResponseBody
// @Failure 400 {object} controllers.SetDocumentStatus.invalidResponseBody
// @Failure 500 {object} controllers.SetDocumentStatus.failedResponseBody
// @Router /document/update/status [put]
func SetDocumentStatus(c *gin.Context) {
	type requestBody struct {
		DocumentID uint   `json:"document_id" binding:"required" example:"1"`
		AppoverID  uint   `json:"approver_id" example:"1"`
		Status     string `json:"status" binding:"required" example:"REJECT"`
		Commnet    string `json:"comment" binding:"required" example:"It looks bad :("`
	}
	type invalidResponseBody struct {
		Error string `json:"error" example:"Invalid request body"`
	}
	type failedResponseBody struct {
		Error string `json:"error" example:"Failed to set document status"`
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

	err = services.SetDocumentStatus(body.DocumentID, body.Status, body.AppoverID, body.Commnet)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, failedResponseBody{
			Error: "Failed to set document status",
		})
		return
	}

	c.JSON(http.StatusOK, successResponseBody{})
}

type historyDto struct {
	Status    string `json:"status" example:"EDIT"`
	Comment   string `json:"comment" example:"It looks bad :("`
	CreatedAt string `json:"created_at" example:"2021-08-01T00:00:00Z"`
}

// @Summary Get document histories
// @Description Get all histories of the document; the document has to exist, and the user has to have permission to view such document.
// @Tags Document
// @Accept json
// @Produce json
// @Param body body controllers.GetDocumentHistories.requestBody true " "
// @Success 200 {object} controllers.GetDocumentHistories.successResponseBody
// @Failure 400 {object} controllers.GetDocumentHistories.invalidResponseBody
// @Failure 500 {object} controllers.GetDocumentHistories.failedResponseBody
// @Router /document/histories [post]
func GetDocumentHistories(c *gin.Context) {
	type requestBody struct {
		DocumentID uint `json:"document_id" binding:"required" example:"1"`
		UserID     uint `json:"user_id" binding:"required" example:"1"`
	}
	type invalidResponseBody struct {
		Error string `json:"error" example:"Invalid request body"`
	}
	type failedResponseBody struct {
		Error string `json:"error" example:"Failed to get document history"`
	}
	type successResponseBody struct {
		Histories []historyDto `json:"histories"`
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

	histories, err := services.GetDocumentHistories(body.DocumentID, body.UserID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, failedResponseBody{
			Error: "Failed to get document history",
		})
		return
	}

	var historiesDto []historyDto

	for _, history := range histories {
		historiesDto = append(historiesDto, historyDto{
			Status:    history.Status,
			Comment:   history.Comment,
			CreatedAt: history.CreatedAt.String(),
		})
	}

	c.JSON(http.StatusOK, successResponseBody{
		Histories: historiesDto,
	})
}

// @Summary Delete document
// @Description Delete the document; the user should have authorization to delete.
// @Tags Document
// @Accept json
// @Produce json
// @Param body body controllers.DeleteDocument.requestBody true " "
// @Success 200 {object} controllers.DeleteDocument.successResponseBody
// @Failure 400 {object} controllers.DeleteDocument.invalidResponseBody
// @Failure 500 {object} controllers.DeleteDocument.failedResponseBody
// @Router /document [delete]
func DeleteDocument(c *gin.Context) {
	type requestBody struct {
		AuthorID   uint `json:"author_id" binding:"required" example:"1"`
		DocumentID uint `json:"document_id" binding:"required" example:"1"`
	}
	type invalidResponseBody struct {
		Error string `json:"error" example:"Invalid request body"`
	}
	type failedResponseBody struct {
		Error string `json:"error" example:"Failed to delete document"`
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

	err = services.DeleteDocument(body.AuthorID, body.DocumentID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, failedResponseBody{
			Error: "Failed to delete document",
		})
		return
	}

	c.JSON(http.StatusOK, successResponseBody{})
}

// @Summary Get document content
// @Description Get the document content; the user should have authorization to get the content.
// @Tags Document
// @Accept json
// @Produce json
// @Param body body controllers.GetDocumentContent.requestBody true " "
// @Success 200 {object} controllers.GetDocumentContent.successResponseBody
// @Failure 400 {object} controllers.GetDocumentContent.invalidResponseBody
// @Failure 500 {object} controllers.GetDocumentContent.failedResponseBody
// @Router /document/content [post]
func GetDocumentContent(c *gin.Context) {
	type requestBody struct {
		UserID     uint `json:"user_id" binding:"required" example:"1"`
		DocumentID uint `json:"document_id" binding:"required" example:"1"`
	}
	type invalidResponseBody struct {
		Error string `json:"error" example:"Invalid request body"`
	}
	type failedResponseBody struct {
		Error string `json:"error" example:"Failed to get document content"`
	}
	type successResponseBody struct {
		Document models.Document `json:"document"`
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

	document, err := services.GetDocumentContent(body.UserID, body.DocumentID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, failedResponseBody{
			Error: "Failed to get document content",
		})
		return
	}

	c.JSON(http.StatusOK, successResponseBody{
		Document: document,
	})
}
