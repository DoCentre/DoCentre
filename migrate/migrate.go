package main

import (
	"github.com/docentre/docentre/initializers"
	"github.com/docentre/docentre/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// auto migrate the user model (create users table in database)
	initializers.DB.AutoMigrate(&models.User{})
}
