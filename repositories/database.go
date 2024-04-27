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

// MigrateDB creates tables, missing foreign keys, constraints, columns and indexes.
// It will change existing column's type if its size, precision, nullable changed.
// It WON’T delete unused columns to protect your data.
func MigrateDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Document{})
}
