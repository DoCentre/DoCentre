package services

import (
	"fmt"

	"github.com/docentre/docentre/models"
	"github.com/docentre/docentre/repositories"
	"gorm.io/gorm"
)

func CreateDocument(authorID uint) (models.Document, error) {
	doc := models.Document{AuthorID: authorID}
	result := repositories.DB.Omit("ApprovedDate, ApproverID, ViewerID").Create(&doc)
	if result.Error != nil {
		return models.Document{}, result.Error
	}

	return doc, nil
}

type EmailSendingError struct {
	Err error
}

func (e EmailSendingError) Error() string {
	return e.Err.Error()
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

// UpdateDocument an email is sent to the approver if the status is "APPROVE" and returns an `EmailSendingError` if the email sending fails.
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

	result = repositories.DB.Model(&models.DocumentHistory{}).Create(&models.DocumentHistory{
		DocumentID: doc.DocumentID,
		Status:     doc.Status,
	})
	if result.Error != nil {
		return 0, result.Error
	}

	// NOTE: Put at the end so that an email sending failure does not affect the status change.
	if doc.Status == "APPROVE" {
		if err := sendEmailToApprover(doc.DocumentID, doc.ApproverID); err != nil {
			return 0, EmailSendingError{err}
		}
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

func GetApproverDocuments(approverID uint) ([]models.Document, error) {
	var docs []models.Document
	result := repositories.DB.Where("approver_id = ?", approverID).Find(&docs)
	if result.Error != nil {
		return []models.Document{}, result.Error
	}
	return docs, nil
}

func GetVerifyDocuments(viewerID uint) ([]models.Document, error) {
	var docs []models.Document
	var result *gorm.DB
	if isAdmin(viewerID) {
		result = repositories.DB.Find(&docs)
	} else {
		result = repositories.DB.Where("status = ?", "VERIFY").Find(&docs)
	}
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

func GetDocumentContent(userID, documentID uint) (models.Document, error) {
	var document models.Document
	result := repositories.DB.Where("id = ?", documentID).First(&document)
	if result.Error != nil {
		return models.Document{}, result.Error
	}
	return document, nil
}
