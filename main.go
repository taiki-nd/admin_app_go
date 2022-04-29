package main

import (
	"admin_app_go/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println(config.Config.SqlDevelop)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
