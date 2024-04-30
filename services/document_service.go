package services

import (
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
