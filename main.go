package main

import (
	"admin_app_go/config"
	"admin_app_go/db"
	"admin_app_go/routes"
	"admin_app_go/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	utils.Logging(config.Config.Logfile)
	db.ConnectToDb()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	})) //他portからのアクセスを有効にする為

	routes.Routes(app)

	app.Listen(":3000")
}
