package services

import (
	"time"

	"github.com/docentre/docentre/models"
	"github.com/docentre/docentre/repositories"
)

func CreateDocument(authorID uint) (models.Document, error) {
	doc := models.Document{AuthorID: authorID, LastEditDate: time.Now()}
	result := repositories.DB.Create(&doc)
	if result.Error != nil {
		return models.Document{}, result.Error
	}

	return doc, nil
}
