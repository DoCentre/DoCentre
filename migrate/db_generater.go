package migrate

import (
	"github.com/docentre/docentre/models"
	"github.com/docentre/docentre/repositories"
)

// CreateDatabase creates the tables used in this application.
func CreateDatabase() {
	repositories.DB.AutoMigrate(&models.User{})
	repositories.DB.AutoMigrate(&models.Document{})
}
