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

func UpdateDocument(documentID uint, authorID uint, title string, content string, appendix string, status string, approverID uint) (uint, error) {

	if approverID == 0 {
		result := repositories.DB.Model(&models.Document{}).Where("id = ?", documentID).Updates(map[string]interface{}{"author_id": authorID, "title": title, "content": content, "appendix": appendix, "status": status})
		
		if result.Error != nil {
			return 0, result.Error
		}
		return documentID, nil

	} else {
		result := repositories.DB.Model(&models.Document{}).Where("id = ?", documentID).Updates(map[string]interface{}{"author_id": authorID, "title": title, "content": content, "appendix": appendix, "status": status, "approver_id": approverID})

		if result.Error != nil {
			return 0, result.Error
		}
		return documentID, nil
	}

}
