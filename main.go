package main

import (
	"github.com/docentre/docentre/router"

	"github.com/docentre/docentre/repositories"
)

func init() {
	repositories.LoadEnvVariables()
	repositories.ConnectToDB()
	repositories.CreateDatabase()
}

func main() {
	r := router.SetupRouter()

	r.Run()
}
