package main

import (
	"github.com/docentre/docentre/migrate"
	"github.com/docentre/docentre/router"

	"github.com/docentre/docentre/repositories"
)

func init() {
	repositories.LoadEnvVariables()
	repositories.ConnectToDB()
	migrate.CreateDatabase()
}

func main() {
	r := router.SetupRouter()

	r.Run()
}
