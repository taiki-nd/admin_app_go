package main

import (
	"admin_app_go/config"
	"admin_app_go/db"
	"admin_app_go/routes"
	"admin_app_go/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	utils.Logging(config.Config.Logfile)
	db.ConnectToDb()

	app := fiber.New()

	routes.Routes(app)

	app.Listen(":3000")
}
