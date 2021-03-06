package main

import (
	"github.com/cciobanu98/donation/app"
	"github.com/cciobanu98/donation/config"
)

// @title Swagger Project API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api

func main() {
	config := config.LoadConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}