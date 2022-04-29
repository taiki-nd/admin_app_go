package main

import (
	"admin_app_go/config"
	"admin_app_go/utils"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println(config.Config.SqlDevelop)

	utils.Logging(config.Config.Logfile)
	log.Println("logging test")
	log.Fatalln("stop test")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
