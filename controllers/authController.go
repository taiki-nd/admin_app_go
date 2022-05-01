package controllers

import (
	"admin_app_go/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	user := models.User{
		FirstName: "Taiki",
		LastName:  "Noda",
		Email:     "a@a.com",
		Password:  "xxxxxxxx",
	}

	return c.JSON(user)
}
