package repositories

import (
	"log"

	"github.com/docentre/docentre/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectToDB connects to the database and stores the connection in the DB package variable.
func ConnectToDB(dialector gorm.Dialector, config *gorm.Config) {
	var err error
	DB, err = gorm.Open(dialector, config)
	if err != nil {
		log.Fatal("failed to connect database")
	}
}

// MigrateDB creates tables, missing foreign keys, constraints, columns and indexes.
// It will change existing column's type if its size, precision, nullable changed.
// It WON’T delete unused columns to protect your data.
func MigrateDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Document{})
	DB.AutoMigrate(&models.DocumentViewer{})
	DB.AutoMigrate(&models.DocumentHistory{})
}
