package main

import (
	"admin_app_go/config"
	"admin_app_go/db"
	"admin_app_go/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	utils.Logging(config.Config.Logfile)
	db.ConnectToDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
