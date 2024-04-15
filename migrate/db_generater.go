package migrate

import (
	"github.com/docentre/docentre/initializers"
	"github.com/docentre/docentre/models"
)

// CreateDatabase creates the tables used in this application.
func CreateDatabase() {
	
	initializers.DB.AutoMigrate(&models.User{})

}
