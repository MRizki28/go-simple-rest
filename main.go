package main

import (
	"github.com/MRizki28/go-simple-rest/src/config"
	"github.com/MRizki28/go-simple-rest/src/routes"
)

func main() {
	r := routes.SetupRoutes()
	config.DatabaseConnetion()
	r.Run(":8890")
}