package main

import (
	"os"

	"github.com/docentre/docentre/router"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/docentre/docentre/repositories"
)

func init() {
	dsn := os.Getenv("DB_URL")
	// NOTE: Decouple the kind of database from the connection,
	// so that multiple databases can be supported.
	repositories.ConnectToDB(mysql.Open(dsn), &gorm.Config{})
	repositories.MigrateDB()
}

func main() {
	r := router.SetupRouter()

	r.Run()
}
