package controllers

import "github.com/gofiber/fiber/v2"

func Bye(c *fiber.Ctx) error {
	return c.SendString("Bye, World 👋!")
}
