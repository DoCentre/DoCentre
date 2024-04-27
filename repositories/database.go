package repositories

import (
	"log"
	"os"

	"github.com/docentre/docentre/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
}

// CreateDatabase creates the tables used in this application.
func CreateDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Document{})
}
