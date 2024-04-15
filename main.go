package main

import (
	"github.com/docentre/docentre/migrate"
	"github.com/docentre/docentre/router"

	"github.com/docentre/docentre/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrate.CreateDatabase()
}

func main() {
	r := router.SetupRouter()

	r.Run()
}
