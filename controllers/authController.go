package controllers

import (
	"admin_app_go/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("POST method error: %s", err)
		return err
	}

	log.Println(data)

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password & password_confirm dose not match.",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  data["password"],
	}

	return c.JSON(user)
}
