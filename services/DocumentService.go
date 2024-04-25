package services

import (
	"time"

	"github.com/docentre/docentre/initializers"
	"github.com/docentre/docentre/models"
)

func CreateDocument(authorID uint) (models.Document, error) {
	doc := models.Document{AuthorID: authorID, LastEditDate: time.Now()}
	result := initializers.DB.Create(&doc)
	if result.Error != nil {
		return models.Document{}, result.Error
	}

	return doc, nil
}
