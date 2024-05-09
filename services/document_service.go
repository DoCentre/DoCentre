package services

import (
	"fmt"

	"github.com/docentre/docentre/models"
	"github.com/docentre/docentre/repositories"
)

func CreateDocument(authorID uint) (models.Document, error) {
	doc := models.Document{AuthorID: authorID}
	result := repositories.DB.Omit("ApprovedDate, ApproverID, ViewerID").Create(&doc)
	if result.Error != nil {
		return models.Document{}, result.Error
	}

	return doc, nil
}

type UpdateDocumentSnapshot struct {
	DocumentID uint
	AuthorID   uint
	Title      string
	Content    string
	Appendix   string
	Status     string
	ApproverID uint
}

func UpdateDocument(doc UpdateDocumentSnapshot) (uint, error) {
	// documentID uint, authorID uint, title string, content string, appendix string, status string, approverID uint
	var noApproverYet uint = 0
	var updateString map[string]interface{}

	if doc.ApproverID == noApproverYet {
		updateString = map[string]interface{}{"author_id": doc.AuthorID, "title": doc.Title, "content": doc.Content, "appendix": doc.Appendix, "status": doc.Status}
	} else {
		updateString = map[string]interface{}{"author_id": doc.AuthorID, "title": doc.Title, "content": doc.Content, "appendix": doc.Appendix, "status": doc.Status, "approver_id": doc.ApproverID}
	}

	result := repositories.DB.Model(&models.Document{}).Where("id = ?", doc.DocumentID).Updates(updateString)

	if result.Error != nil {
		return 0, result.Error
	}
	return doc.DocumentID, nil
}

func GetAuthorDocuments(authorID uint) ([]models.Document, error) {
	var docs []models.Document

	result := repositories.DB.Where("author_id = ?", authorID).Find(&docs)

	if result.Error != nil {
		return []models.Document{}, result.Error
	}

	return docs, nil
}

func GetViewerDocuments(viewerID uint) ([]models.Document, error) {
	var docs []models.Document

	result := repositories.DB.Joins("JOIN document_viewers ON documents.id = document_viewers.document_id").Where("viewer_id = ?", viewerID).Find(&docs)

	if result.Error != nil {
		return []models.Document{}, result.Error
	}

	return docs, nil
}

func AddViewer(documentID, viewerID uint) error {
	docViewer := models.DocumentViewer{DocumentID: documentID, ViewerID: viewerID}
	result := repositories.DB.Create(&docViewer)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// SetDocumentStatus sets the status of the document, as well as the approverID (if present) and comment.
// These changes are also recorded in the DocumentHistory table.
func SetDocumentStatus(documentID uint, status string, approverID uint, comment string) error {
	// Set the status of the document
	var noApproverYet uint = 0
	var statusMap map[string]interface{}
	if approverID == noApproverYet {
		statusMap = map[string]interface{}{"status": status}
	} else {
		statusMap = map[string]interface{}{"status": status, "approver_id": approverID}
	}
	result := repositories.DB.Model(&models.Document{}).Where("id = ?", documentID).Updates(statusMap)
	if result.Error != nil {
		return result.Error
	}

	// Record the status change in the DocumentHistory table.
	result = repositories.DB.Model(&models.DocumentHistory{}).Create(&models.DocumentHistory{
		DocumentID: documentID,
		Status:     status,
		Comment:    comment,
	})
	return result.Error
}

// GetDocumentHistories returns the history of the document with the given documentID. The user indicated by userID has to have permission to view the document.
// In case that the user does not have permission, an error is returned.
func GetDocumentHistories(documentID uint, userID uint) ([]models.DocumentHistory, error) {
	var document models.Document
	result := repositories.DB.Where("id = ?", documentID).First(&document)
	if result.Error != nil {
		return []models.DocumentHistory{}, result.Error
	}

	// Check if the user has permission to view the document.
	if document.AuthorID != userID && document.ApproverID != userID && !isAdmin(userID) {
		return []models.DocumentHistory{}, fmt.Errorf("user %d does not have the permission to view the document %d", userID, documentID)
	}

	var histories []models.DocumentHistory
	result = repositories.DB.Where("document_id = ?", documentID).Find(&histories)
	if result.Error != nil {
		return []models.DocumentHistory{}, result.Error
	}
	return histories, nil
}

func isAdmin(userID uint) bool {
	var user models.User
	result := repositories.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return false
	}
	return user.Identity == "admin"
}

func DeleteDocument(authorID, documentID uint) error {
	var document models.Document
	result := repositories.DB.Where("id = ?", documentID).First(&document)
	if result.Error != nil {
		return fmt.Errorf("document %d not found", documentID)
	}

	// Check if the user has permission to delete the document.
	if document.AuthorID != authorID && !isAdmin(authorID) {
		return fmt.Errorf("user %d does not have the permission to delete the document %d", authorID, documentID)
	}

	result = repositories.DB.Delete(&document)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

type DocumentContent struct {
	Title   string
	Content string
}

func GetDocumentContent(userID, documentID uint) (DocumentContent, error) {
	var document models.Document
	result := repositories.DB.Where("id = ?", documentID).First(&document)
	if result.Error != nil {
		return DocumentContent{}, result.Error
	}

	// Check if the user has permission to get the document content.
	if document.AuthorID != userID && document.ApproverID != userID && !isAdmin(userID) {
		return DocumentContent{}, fmt.Errorf("user %d does not have the permission to view the document %d", userID, documentID)
	}

	return DocumentContent{Title: document.Title, Content: document.Content}, nil
}
