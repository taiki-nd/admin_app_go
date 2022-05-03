package controllers

import (
	"admin_app_go/db"
	"admin_app_go/models"

	"github.com/gofiber/fiber/v2"
)

func UserIndex(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Find(&users)

	return c.JSON(users)
}
