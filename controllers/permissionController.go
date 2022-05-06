package controllers

import (
	"admin_app_go/db"
	"admin_app_go/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func PermissionIndex(c *fiber.Ctx) error {
	var permissions []models.Permission
	db.DB.Find(&permissions)

	log.Println("show all users")

	return c.JSON(permissions)
}
